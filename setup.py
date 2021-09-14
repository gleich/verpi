import shutil
import os
import sys
import time
from urllib.request import urlopen

tmp_dir = "/home/pi/verpi_build"
go_binary_path = os.path.join(tmp_dir, "bin/go")
systemd_service_path = "/etc/systemd/system/verpi.service"
verpi_bin_path = "/usr/local/bin/verpi"


def main() -> None:
    command = sys.argv[1]
    if command == "install":
        setup()
        install_go()
        clone_repo()
        compile()
        install_verpi()
        install_verpi_service()
    elif command == "uninstall":
        uninstall()
    else:
        print(command, "isn't a valid command")
        exit(1)
    os.chdir("..")
    shutil.rmtree(tmp_dir)
    reboot()


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
    print("Installing verpi at", verpi_bin_path)
    if os.path.exists(verpi_bin_path):
        os.remove(verpi_bin_path)
    os.rename("verpi/dist/verpi", verpi_bin_path)
    print("verpi installed at", verpi_bin_path)


def install_verpi_service() -> None:
    print("Installing systemd service for verpi")
    with urlopen(
        "https://raw.githubusercontent.com/gleich/verpi/master/verpi.service"
    ) as response:
        content = response.read().decode("utf-8")
    with open(systemd_service_path, "w") as systemd_file:
        systemd_file.write(content)
    command("systemctl enable verpi")
    print("Added and started systemd service")


def uninstall() -> None:
    os.remove(verpi_bin_path)
    print("Deleted binary")
    command("systemctl disable verpi")
    print("Disabled systemd service")
    os.remove(systemd_service_path)
    print("Deleted systemd service file")


def reboot() -> None:
    print(
        "\nRebooting pi in 3 seconds. You might need to cut the power to fully turn off the lights."
    )
    for i in reversed(range(2)):
        print(i + 1)
        time.sleep(1)
    print("Rebooting now")
    command("reboot")


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
