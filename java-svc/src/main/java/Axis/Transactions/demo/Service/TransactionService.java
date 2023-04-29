package Axis.Transactions.demo.Service;

import Axis.Transactions.demo.Model.TransactionModel;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class TransactionService {

    @Autowired
    private MongoTemplate mongoTemplate;


    public TransactionModel getTransactionById(String id) {
        Query query = new Query(Criteria.where("txnDetailsId").is(id));
        return mongoTemplate.findOne(query, TransactionModel.class);
    }


    public List<TransactionModel> getTransactionPaginated(int rpi, int rps) {
        Pageable pageable = PageRequest.of(rpi, rps);
        Query query = new Query().with(pageable);
        return mongoTemplate.find(query, TransactionModel.class);
    }

}
