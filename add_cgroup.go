package main 
import  (
	"os/exec" 
	"path" 
	"os"
	"fmt"
	"io/ioutil"
	"syscall"
	"strconv"
)
const Mount = "/sys/fs/cgroup/memory"

func main() {
	if os.Args[0] == "/proc/self/exe" {
		fmt.Println("current pid:",syscall.Getpid())
		fmt.Println()	
		cmd := exec.Command("sh","-c",`stress --vm-bytes 25m --vm-keep -m 1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{
		}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout 
		cmd.Stderr = os.Stderr 
		check(cmd.Run())
	}  
		cmd := exec.Command("/proc/self/exe")
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Cloneflags : syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS , 
		}
		cmd.Stdin = os.Stdin 
		cmd.Stdout = os.Stdout 
		cmd.Stderr = os.Stderr
		check(cmd.Start())
		fmt.Printf("%v\n",cmd.Process.Pid) // 新建的进程在外部命名空间中的pid 
		os.Mkdir(path.Join(Mount,"testmemorylimit"),0755) // 创建cgroup
		ioutil.WriteFile(path.Join(Mount,"testmemorylimit","tasks"),[]byte(strconv.Itoa(cmd.Process.Pid)),0777) //  将容器进程写入task，就是加入这个cgroup
		ioutil.WriteFile(path.Join(Mount,"testmemorylimit","memory.limit_in_bytes"),[]byte("25m"),0777) //  限制内存 
		cmd.Process.Wait()

	}
// 其实我觉得这个程序写的很奇怪。。。。
func check(err error ) {
	if err != nil {
		fmt.Println("Error: ",err) 
		os.Exit(1) 
	}
}