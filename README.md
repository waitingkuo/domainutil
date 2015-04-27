# domainutil
Simple domain parser

## Example

    package main
    
    import (
      "fmt"
      "github.com/waitingkuo/domainutil"
    )
    
    func main() {
    
      domain, err := domainutil.ParseFromRawURL("http://www.google.com")
      if err != nil {
        panic(err)
      }
    
      fmt.Println("RootDomain:", domain.RootDomain)
      fmt.Println("SubDomain:", domain.SubDomain)
    }

### Output

    RootDomain: google.com
    SubDomain: www
    

### Others
Feel free to leave any commands or issues
