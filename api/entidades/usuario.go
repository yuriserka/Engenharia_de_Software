package entidades

// Usuario é o principal ator do sistema.
//
// Qualquer usuário pode obter dados sobre os eventos, adicionar ou editar um evento que ele criou e
// comprar ingressos para outros.
//
// Ao ser cadastrado, pode cadastrar, descadastrar e alterar evento, além de poder retirar sua conta do sistema.
//
// Só é permitido a gerencia de 5 eventos.
type Usuario struct {
	Cpf, Senha string
}
