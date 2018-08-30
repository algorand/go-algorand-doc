package nodecontrol

import (
	"encoding/json"
	"fmt"
	"github.com/algorand/go-algorand/api/algod"
	"github.com/algorand/go-algorand/util"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// NodeController provides an object for controlling a specific algod node instance
type NodeController struct {
	algod         string
	dataDir       string
	pidFile       string
	netFile       string
	netListenFile string
}

// MakeNodeController creates a NodeController representing a
// specific data directory (and an associated binary directory)
func MakeNodeController(binDir, dataDir string) NodeController {
	nc := NodeController{
		algod:         filepath.Join(binDir, "algod"),
		dataDir:       dataDir,
		pidFile:       filepath.Join(dataDir, "algod.pid"),
		netFile:       filepath.Join(dataDir, "algod.net"),
		netListenFile: filepath.Join(dataDir, "algod-listen.net"),
	}

	return nc
}

// NodeStartArgs represents the possible arguments for starting the algod node process.
type NodeStartArgs struct {
	PeerAddress string
	ListenIP    string
}

// Start will start the node if not already started.
// Returns false, nil if started successfully (and not already running)
func (nc NodeController) Start(args NodeStartArgs) (alreadyRunning bool, err error) {
	_, err = nc.Status()
	if err == nil {
		return true, nil
	}

	alreadyRunning = false

	startArgs := make([]string, 0)
	startArgs = append(startArgs, "-d")
	startArgs = append(startArgs, nc.dataDir)
	peerDial := args.PeerAddress
	if len(peerDial) > 0 {
		startArgs = append(startArgs, "-p")
		startArgs = append(startArgs, peerDial)
	}
	listenIP := args.ListenIP
	if len(listenIP) > 0 {
		startArgs = append(startArgs, "-l")
		startArgs = append(startArgs, listenIP)
	}
	subcmd := exec.Command(nc.algod, startArgs...)
	subcmd.Stderr = os.Stderr
	err = subcmd.Start()
	// TODO we must collect stdout and stderr to catch panics
	// TODO we should designate a new (unique) log file
	if err != nil {
		return
	}

	// Wait on the algod process and check if exits
	c := make(chan bool)
	go func() {
		subcmd.Wait()
		c <- true
	}()

	success := false
	for !success {
		select {
		case <-c:
			return false, fmt.Errorf("node exited before we could contact it")
		case <-time.After(time.Second):
			_, err := nc.Status()
			if err == nil {
				success = true
			}
		}
	}
	return
}

// Stop determines the node's PID from its PID file and uses that to kill it.
func (nc NodeController) Stop() error {
	pid, err := nc.GetPID()
	if err != nil {
		return err
	}
	return syscall.Kill(int(pid), syscall.SIGTERM)
}

// Status retrieves the StatusResponse from the running node
func (nc NodeController) Status() (response algod.StatusResponse, err error) {
	err = nc.Get(&response, "/status", nil)
	return
}

func (nc NodeController) serverURL() (url.URL, error) {
	addr, err := nc.GetHostAddress()
	if err != nil {
		return url.URL{}, err
	}
	return url.URL{Scheme: "http", Host: addr}, nil
}

// GetHostAddress retrieves the REST address for the node from its algod.net file.
func (nc NodeController) GetHostAddress() (string, error) {
	// For now, we want the old behavior to 'just work';
	// so if data directory is not specified, we assume the default address of 127.0.0.1:8080
	if len(nc.dataDir) == 0 {
		return "127.0.0.1:8080", nil
	}
	return util.GetFirstLineFromFile(nc.netFile)
}

// GetListeningAddress retrieves the listening address from the algod-listen.net file for the node
func (nc NodeController) GetListeningAddress() (string, error) {
	return util.GetFirstLineFromFile(nc.netListenFile)
}

// extractError checks if the response signifies an error (for now, StatusCode != 200).
// If so, it returns the error.
// Otherwise, it returns nil.
func extractError(resp *http.Response) error {
	if resp.StatusCode == 200 {
		return nil
	}

	var errorBuf [algod.ErrorMaxSize]byte
	resp.Body.Read(errorBuf[:]) // ignore returned error
	return fmt.Errorf("HTTP %v: %s", resp.Status, errorBuf)
}

// Get performs a GET request to the specific path against the node
// TODO add query parameters as arguments? or put into a getQuery function
func (nc NodeController) Get(response interface{}, path string, request interface{}) error {
	queryURL, err := nc.serverURL()
	if err != nil {
		return err
	}
	queryURL.Path = path

	if request != nil {
		v, err := query.Values(request)
		if err != nil {
			return err
		}

		queryURL.RawQuery = v.Encode()
	}

	resp, err := http.Get(queryURL.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = extractError(resp)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(resp.Body)
	return dec.Decode(&response)
}

// PostQuery sends a POST request to the given path with the given request object.
// No query parameters will be sent if request is nil.
// response must be a pointer to an object as postQuery writes the response there.
func (nc NodeController) PostQuery(response interface{}, path string, request interface{}) error {
	queryURL, err := nc.serverURL()
	if err != nil {
		return err
	}
	queryURL.Path = path

	if request != nil {
		v, err := query.Values(request)
		if err != nil {
			return err
		}

		queryURL.RawQuery = v.Encode()
	}

	resp, err := http.PostForm(queryURL.String(), url.Values{})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = extractError(resp)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(resp.Body)
	return dec.Decode(&response)
}

// GetPID returns the PID from the algod.pid file in the node's data directory, or an error
func (nc NodeController) GetPID() (pid uint64, err error) {
	pidStr, err := ioutil.ReadFile(nc.pidFile)
	if err != nil {
		return 0, err
	}

	pid, err = strconv.ParseUint(strings.TrimSuffix(string(pidStr), "\n"), 10, 32)
	return
}
