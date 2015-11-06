package edu.virginia.depositauthws.models;

import com.fasterxml.jackson.annotation.JsonProperty;
import org.hibernate.validator.constraints.Length;

public class AllStatus {
    private long id;

    @Length(max = 3)
    private String content;

    public AllStatus() {
        // Jackson deserialization
    }

    public AllStatus(long id, String content) {
        this.id = id;
        this.content = content;
    }

    @JsonProperty
    public long getId() {
        return id;
    }

    @JsonProperty
    public String getContent() {
        return content;
    }
}