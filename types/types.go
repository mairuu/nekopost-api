package types

type Category struct {
	CateName  string
	CateLink  string
	CateCode  string
}

type ChapterApi interface {
    GetChapters(page int, _type string) ([]byte, error)
}

type CommentApi interface {
    GetComments(id int) ([]byte, error)
}

type ProjectApi interface {
    GetProject(pid int) ([]byte, error)
    GetProjects(page int, order string, types, genres []string) ([]byte, error)
    GetRandomProjects() ([]byte, error)
    GetPopularProjects() ([]byte, error)
    GetProjectChapter(pid, cid int) ([]byte, error)
}
