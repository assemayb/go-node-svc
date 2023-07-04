package Axis.Transactions.demo.Controller;


import Axis.Transactions.demo.Model.TransactionModel;
import Axis.Transactions.demo.Service.TransactionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;


@RestController
@RequestMapping("/transaction")
public class TransactionController {

    @Autowired
    private TransactionService transactionService;

    @GetMapping("/{id}")
    public TransactionModel getTransactionById( @PathVariable String id) {
        return transactionService.getTransactionById(id);
    }

    @GetMapping("/txns")
    public List<TransactionModel> getTxnsPaginated(@RequestParam("rpi") int rpi, @RequestParam("rps") int rps) {
        return transactionService.getTransactionPaginated(rpi, rps);
    }
}
