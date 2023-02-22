namespace ex.vsa.GetAccount;

using scyna;
using ex.vsa.model;
using ex.vsa.repository;

public class Handler : Endpoint.Handler<proto.RegisterAccountRequest>
{
    public override void Execute()
    {
        context.Info("Receive CreateAccountRequest");
        var repository = BaseRepository.load(context);
        var account = repository.GetAccountByEmail(request.Email);
        Response(new proto.GetAccountResponse
        {
            ID = account.ID,
            Name = account.Name,
            Email = account.Email,
        });
    }
}