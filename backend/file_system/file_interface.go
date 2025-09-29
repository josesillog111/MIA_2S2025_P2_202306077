package fsys

type FileSystem interface {
	Mkfs(id string, typeFormat string) error
	Cat(paths ...string) (string, error)
	Login(user string, pass string) error
	Logout() error
	Mkgrp(name string) error
	Rmgrp(name string) error
	Mkusr(name string, pass string, grp string) error
	Rmusr(name string) error
	Chgrp(name string, newgrp string) error
	Mkfile(virtualPath string, recursive bool, size int64, content string) error
	Mkdir(path string, p bool) error
	Rep(name string, path string, path_file_ls string, format string) error
}
