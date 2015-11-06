package edu.virginia.depositauthws.health;

import com.codahale.metrics.health.HealthCheck;

public class ServiceHealthCheck extends HealthCheck {
    private final String dirname;

    public ServiceHealthCheck(String dirname) {
        this.dirname = dirname;
    }

    @Override
    protected Result check() throws Exception {
        //    return Result.unhealthy("template doesn't include a name");
        return Result.healthy();
    }
}