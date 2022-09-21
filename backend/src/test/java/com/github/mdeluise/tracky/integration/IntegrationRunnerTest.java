package com.github.mdeluise.tracky.integration;

import io.cucumber.junit.Cucumber;
import io.cucumber.junit.CucumberOptions;
import org.junit.runner.RunWith;

@RunWith(Cucumber.class)
@CucumberOptions(
    features = "classpath:features",
    glue = {
        "com.github.mdeluise.tracky.integration",
        "com.github.mdeluise.tracky.integration.steps"
    },
    plugin = {"pretty"}
)
public class IntegrationRunnerTest {
}
