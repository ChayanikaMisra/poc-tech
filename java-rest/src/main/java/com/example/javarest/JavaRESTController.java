package com.example.javarest;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;


@RestController
public class JavaRESTController {

    @GetMapping("/hello")
    public String hello(){
        return "Hello";
    }
    
}
