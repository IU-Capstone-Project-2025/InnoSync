package com.innosync.controller;

import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("api/user")
@Tag(name = "User API", description = "API for user authentication") // Swagger annotation
public class UserController {

    @GetMapping("/hello")
    public String hello() {
        return "hello";
    }
}
