module github.com/jojahn/kpgo/10_Enterprise_Programming/client

go 1.15

require github.com/jojahn/kpgo/10_Enterprise_Programming/mail v1.0.0
replace github.com/jojahn/kpgo/10_Enterprise_Programming/mail v1.0.0 => ../mail

require github.com/jojahn/kpgo/10_Enterprise_Programming/mail/v2 v2.0.0
replace github.com/jojahn/kpgo/10_Enterprise_Programming/mail/v2 v2.0.0 => ../mail/v2
