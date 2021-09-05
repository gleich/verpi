import shutil
import os
import sys

tmp_dir = "/home/pi/verpi_build"
go_binary_path = os.path.join(tmp_dir, "bin/go")


def main() -> None:
    command = sys.argv[1]
    if command == "install":
        setup()
        install_go()
        clone_repo()
        compile()
        install_verpi()
    os.chdir("..")
    shutil.rmtree(tmp_dir)


def setup() -> None:
    if os.path.exists(tmp_dir):
        shutil.rmtree(tmp_dir)
    os.mkdir(tmp_dir)
    os.chdir(tmp_dir)


def install_go() -> None:
    go_version = "1.17"
    tar_file = f"go{go_version}.linux-armv6l.tar.gz"
    print(f"Installing temporaray version of go {go_version}...")
    command("wget -c https://golang.org/dl/" + tar_file)
    command(f"tar -C {tmp_dir} -xvzf {tar_file}")
    print(f"Setup temporaray version of go {go_version}")


def clone_repo() -> None:
    print("Cloning repo")
    command("git clone https://github.com/gleich/verpi.git")
    print("Cloned repo")


def compile() -> None:
    print("Compiling binary from source code:")
    original_gopath = os.getenv("GOPATH")
    os.environ["GOPATH"] = os.path.join(tmp_dir, "goroot")
    os.chdir("verpi")
    command("../go/bin/go build -v -o dist/verpi .")
    if original_gopath is not None:
        os.environ["GOPATH"] = original_gopath
    os.chdir("..")
    print("Compiled binary")


def install_verpi() -> None:
    verpi_bin_path = "/usr/local/bin/verpi"
    print("Installing verpi at", verpi_bin_path)
    if os.path.exists(verpi_bin_path):
        os.remove(verpi_bin_path)
    os.rename("verpi/dist/verpi", verpi_bin_path)
    print("verpi installed at", verpi_bin_path)


def command(cmd: str) -> None:
    """Run os.system but exit with a failure if the exit code is not zero

    Args:
        cmd (str): Command to run
    """
    code = os.system(cmd)
    if code != 0:
        print(f"Failed to run {cmd} with status code of {code}")
        exit(1)


if __name__ == "__main__":
    main()
