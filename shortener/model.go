package shortener

type RedirectService interace {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}