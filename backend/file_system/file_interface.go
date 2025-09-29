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
	Remove(path string) error
	Edit(path string, contenido string) error
	Rename(path string, name string) error
	Copy(path string, dest string) error
	Move(path string, dest string) error
	Find(path string, name string) (string, error)
	Chown(path string, user string, recursive bool) error
	Chmod(path string, ugo int64, recursive bool) error
	Recovery() error
	Loss() error
	Journaling() error
	Rep(name string, path string, path_file_ls string, format string) error
}
