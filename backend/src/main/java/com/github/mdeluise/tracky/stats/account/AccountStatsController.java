package com.github.mdeluise.tracky.stats.account;

import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/account-stats")
@Tag(name = "Account statistics", description = "Endpoints for retrieving the account statistics")
public class AccountStatsController {
    private final AccountStatsCreator accountStatsCreator;


    @Autowired
    public AccountStatsController(AccountStatsCreator accountStatsCreator) {
        this.accountStatsCreator = accountStatsCreator;
    }


    @GetMapping
    public ResponseEntity<AccountStats> get() {
        return ResponseEntity.ok(accountStatsCreator.getData());
    }
}
