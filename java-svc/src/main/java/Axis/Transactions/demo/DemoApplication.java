package Axis.Transactions.demo;

import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;



@SpringBootApplication
public class DemoApplication {
    public static void main(String[] args) {

        SpringApplication app = new SpringApplication(DemoApplication.class);
        MongoClient mongoClient = MongoClients.create("mongodb://localhost:27017/axisPayCore");
        app.run(DemoApplication.class, args);
    }

}
