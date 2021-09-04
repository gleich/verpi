import shutil
import os


def main() -> None:
    install_go()
    compile()


go_install_loc = "/usr/local/go/"


def install_go() -> None:
    if shutil.which("go") is None:
        go_version = "1.17"
        tar_fname = f"go{go_version}.linux-armv6l.tar.gz"
        print(f"Installing go {go_version}...")

        if os.path.exists(go_install_loc):
            shutil.rmtree(go_install_loc)
        if os.path.exists(tar_fname):
            os.remove(tar_fname)

        system_check(f"wget https://golang.org/dl/{tar_fname}")
        system_check(f"tar --checkpoint -C /usr/local -xvzf {tar_fname}")
        os.remove(tar_fname)
        print(
            "Installed go",
            go_version,
            "in /usr/local/go/. Binaries are in /usr/local/go/bin/",
        )
    else:
        proceed = (
            input(
                "Looks like go is already installed on your system. You need to a version of go greater than 1.15 for verpi to compile. Are you sure you want to proceed with the installation? (y/n)"
            ).lower()
            == "y"
        )
        if not proceed:
            exit(1)


def compile() -> None:
    print("Compiling binary from source code:")
    system_check(os.path.join(go_install_loc, "bin", "go") + " build -v -o dist/verpi")
    print("Compiled binary")


def system_check(cmd: str) -> None:
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
