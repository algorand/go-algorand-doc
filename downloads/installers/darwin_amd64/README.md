# IMPORTANT PLEASE READ

Hi,

Thanks for your interest in joining our TestNet, we are very excited to have you!

Algorand uses this test environment to measure performance and test new features. We collect metrics and data to help improve the protocol and troubleshoot issues (we will not collect this data on the MainNet unless you explicitly allow it).

By proceeding to join the Algorand TestNet you are acknowledging your consent to the collection and temporary storage for the data outlined below.

<table>
  <tr>
    <td>Location</td>
    <td>Access</td>
    <td>Retention Length</td>
  </tr>
  <tr>
    <td>Private Database</td>
    <td>Algorand Employees & Contractors</td>
    <td>6 months then deleted</td>
  </tr>
</table>

<br/>
<table>
  <tr>
    <td>Data Fields</td>
    <td></td>
  </tr>
  <tr>
    <td>Node Name/Statistics</td>
    <td>Node name / GUID<br/>
Log entries for errors<br/>
Disk space<br/>
Memory use<br/>
CPU utilization<br/>
Network connections<br/>
Running processes<br/>
</td>
  </tr>
  <tr>
    <td>Application Usage</td>
    <td>How often are nodes started<br/>
How long do nodes run once started<br/>
How many hosts are running multiple node instances</td>
  </tr>
  <tr>
    <td>Catchup</td>
    <td>How long does a node take to catchup initially<br/>
How many peers contribute to catchup<br/>
How many failures per catchup attempt<br/>
How long does the node spend catching up, and as percentage of time since last caught up</td>
  </tr>
  <tr>
    <td>Accounts
</td>
    <td>At startup, account PubKeys and Balances (and available keys / valid ranges)</td>
  </tr>
  <tr>
    <td>Crash Recovery</td>
    <td>Process recovery time<br/>
State of recovery</td>
  </tr>
  <tr>
    <td>Consensus performance
</td>
    <td>Per-Round stats<br/>
How long in each round<br/>
How many steps in each round<br/>
How often do we detect a partition<br/>
How long does the partition last</td>
  </tr>
  <tr>
    <td>Block Stats
</td>
    <td>How many transactions per block<br/>
How many distinct proposers & voters<br/>
How many equivocating votes / proposals<br/>
Network stats per-block<br/>
How long does block validation take</td>
  </tr>
  <tr>
    <td>REST API</td>
    <td>How often each given endpoint is called<br/>
How often does each endpoint call fail and what is the failure type<br/>
How many people override the default listening port</td>
  </tr>
  <tr>
    <td>Transaction Pool</td>
    <td>Transaction drop rate and type<br/>
How many invalid transactions do we see<br/>
Max / Average pool depth<br/>
How many transactions in pool when we shut down</td>
  </tr>
  <tr>
    <td>Network stats</td>
    <td>Peer Connections<br/>
Reject / Drop Reason</td>
  </tr>
  <tr>
    <td>Database stats</td>
    <td>Min/Max/Average db.Atomic call<br/>
Database size at startup / shutdown<br/>
Block count at startup</td>
  </tr>
  <tr>
    <td>Update.sh stats</td>
    <td>How often is update.sh run<br/>
How often does update.sh find a new update<br/>
How often does update.sh fail when installing an update<br/>
How often is each channel specified<br/>
What parameters are used with Update.sh (esp. -i, -n, -p)<br/>
How many people are running as a service (update.sh can tell this if systemctl succeeds)</td>
  </tr>
</table>
