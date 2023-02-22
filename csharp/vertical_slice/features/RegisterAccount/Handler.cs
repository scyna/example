namespace ex.vsa.CreateAccount;

using scyna;
using ex.vsa.model;

public class Handler : Endpoint.Handler<proto.RegisterAccountRequest>
{
    public override void Execute()
    {
        context.Info("Receive CreateAccountRequest");

        /*TODO: validate input*/

        var repository = Repository.create(context);

        repository.AssureAccountNotExists(request.Email);

        var account = new Account
        {
            ID = Engine.ID.Next(),
            Name = request.Name,
            Email = request.Email,
            Password = request.Password
        };

        repository.RegisterAccount(account);
        Response(new proto.RegisterAccountResponse { ID = account.ID });
    }
}