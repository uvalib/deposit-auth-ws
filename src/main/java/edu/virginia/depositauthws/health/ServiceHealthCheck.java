package edu.virginia.depositauthws.health;

import com.codahale.metrics.health.HealthCheck;
import edu.virginia.depositauthws.db.DepositAuthDAO;

public class ServiceHealthCheck extends HealthCheck {

    private DepositAuthDAO depositAuthDAO;
    private String dirname;

    public ServiceHealthCheck(DepositAuthDAO depositAuthDAO, String dirname ) {
        this.depositAuthDAO = depositAuthDAO;
        this.dirname = dirname;
    }

    @Override
    protected Result check() throws Exception {
        //    return Result.unhealthy("template doesn't include a name");
        return Result.healthy();
    }
}