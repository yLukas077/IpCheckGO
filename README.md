
## **IPCheck Application**

### **Description**

The IPCheck application is a command line tool written in Go that allows users to search for IP addresses and server names on the internet.

### **Installation**

To install the application, you must have Go installed on your machine. You can get the latest version of Go at [https://golang.org/dl/](https://golang.org/dl/).

Once Go is installed, you can clone the repository and build the application:

```
git clone https://github.com/yLukas077/IpCheckGO
cd IpCheckGO
go build -o ipcheck
```

### **How to Use**

The application has two main commands:

1. `ip` - Searches for IP addresses associated with a host name.
2. `servers` - Searches for server names associated with a host name.

**Usage Examples:**

To search for IP addresses of `amazon.com.br`:

```
./ipcheck ip --host amazon.com.br
```

To search for server names of `amazon.com.br`:

```
./ipcheck servers --host amazon.com.br
```

### **Dependencies**

The application uses the following libraries:

- github.com/cpuguy83/go-md2man/v2 v2.0.2
- github.com/russross/blackfriday/v2 v2.1.0
- github.com/urfave/cli v1.22.14

### **License**

MIT

