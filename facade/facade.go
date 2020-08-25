package facade

import "fmt"

type WalletFacade struct {
	account      Account
	notification Notification
}

func NewWalletFacade() *WalletFacade {
	return &WalletFacade{
		account:      &BankAccount{},
		notification: &EmailNotification{},
	}
}

func (w *WalletFacade) AddMoney(amount int) string {
	w.account.AddMoney(amount)
	return w.notification.Notify(amount)
}

type Account interface {
	AddMoney(int)
}

type BankAccount struct {
	saving int
}

func (b *BankAccount) AddMoney(amount int) {
	b.saving += amount
}

type Notification interface {
	Notify(int) string
}

type EmailNotification struct{}

func (e *EmailNotification) Notify(amount int) string {
	return fmt.Sprintf("Email: %d has been saved to your account", amount)
}
