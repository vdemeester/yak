go-hvkvp: Hyper-V Data Exchange using Go
========================================

[![pipeline status](https://gitlab.com/gbraad/go-hvkvp/badges/master/pipeline.svg)](https://gitlab.com/gbraad/go-hvkvp/commits/master)

This is used by the [Minishift](https://github.com/minishift) project.


## Installation
A [Copr project](https://copr.fedorainfracloud.org/coprs/gbraad/go-hvkvp) is available. Installation can be done as follows:

### CentOS 7 (or older versions of Fedora)
```
$ yum install yum-plugin-copr 
$ yum copr enable gbraad/go-hvkvp
$ yum install go-hvkp
```

### Fedora
```
$ dnf copr enable gbraad/go-hvkvp
$ dnf install go-hvkvp
```

### LiveCD Creator
```
repo --name=base --baseurl=http://mirror.centos.org/centos/7/os/x86_64/
repo --name=updates --baseurl=http://mirror.centos.org/centos/7/updates/x86_64/
repo --name=hvkvp --baseurl=https://copr-be.cloud.fedoraproject.org/results/gbraad/go-hvkvp/epel-7-x86_64/

%packages
@core
go-hvkvp
```

### From source / development
```
$ git clone https://github.com/gbraad/go-hvkvp.git
$ cd go-hvkvp
$ go install -v ./cmd/hvkvp
```

### Build RPM package
```
$ dnf install -y rpm-build go-compilers-golang-compiler
$ rpmbuild -ba hvkvp.spec
```

## Usage

### Prepare and send Key-Value pair
```powershell
$VmMgmt = Get-WmiObject -Namespace root\virtualization\v2 -Class Msvm_VirtualSystemManagementService  
$vm = Get-WmiObject -Namespace root\virtualization\v2 -Class Msvm_ComputerSystem -Filter {ElementName = 'MyVM'}

$kvpDataItem = ([WMIClass][String]::Format("\\{0}\{1}:{2}", $VmMgmt.ClassPath.Server, $VmMgmt.ClassPath.NamespacePath, "Msvm_KvpExchangeDataItem")).CreateInstance()

$kvpDataItem.Name = "IpAddress"
$kvpDataItem.Data = "10.0.75.128"
$kvpDataItem.Source = 0

$VmMgmt.AddKvpItems($Vm, $kvpDataItem.PSBase.GetText(1))
```

#### Note
With `$kvpDataItem.Source = 0` the KVP gets stored as `/var/lib/hyperv/.kvp_pool_0`.


### Receive/Read on the host:

#### All records
```
$ ./hvkvp
Key: IpAddress, Value: 10.0.75.128
```

#### Record by specific key
```
$ ./hvkvp --key Hostname
fedora-vm%
```

#### Note
When dealing with special characters, consider using base64 encoding.

For example, the value of the message might contain:
```
REVWSUNFPWV0aDAKVVNFREhDUD15Cg==
```

This can be received on the host as:
```
$ hvkvp -key PROVISION_NETWORKING | base64 --decode > /var/lib/minishift/networking
$ cat /var/lib/minishift/networking
DEVICE=eth0
USEDHCP=y
```

