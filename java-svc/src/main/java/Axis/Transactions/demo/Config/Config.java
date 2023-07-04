package Axis.Transactions.demo.Config;

import org.springframework.context.annotation.Bean;


public class Config {
    @Bean
    public String getBean() {
        return "Hello World";
    }
}
