package com.github.mdeluise.tracky.observation;

import com.github.mdeluise.tracky.authentication.User;
import com.github.mdeluise.tracky.tracker.Tracker;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.Table;
import jakarta.validation.constraints.NotNull;
import org.hibernate.validator.constraints.Length;

import java.util.Date;
import java.util.Objects;

@Entity
@Table(name = "observations")
public class Observation {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;
    @ManyToOne
    @JoinColumn(name = "tracker_id", nullable = false)
    private Tracker tracker;
    @NotNull
    private Date instant = new Date();
    @Length(max = 100)
    private String note;
    @NotNull
    @Column(name = "observation_value")
    private float value;
    @ManyToOne
    @JoinColumn(name = "user_id", nullable = false)
    private User user;


    public Long getId() {
        return id;
    }


    public void setId(Long id) {
        this.id = id;
    }


    public Tracker getTracker() {
        return tracker;
    }


    public void setTracker(Tracker tracker) {
        this.tracker = tracker;
    }


    public Date getInstant() {
        return instant;
    }


    public void setInstant(Date instant) {
        this.instant = instant;
    }


    public Number getValue() {
        return value;
    }


    public void setValue(Number value) {
        this.value = value.floatValue();
    }


    public String getNote() {
        return note;
    }


    public void setNote(String note) {
        this.note = note;
    }


    public User getUser() {
        return user;
    }


    public void setUser(User user) {
        this.user = user;
    }


    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (o == null || getClass() != o.getClass()) {
            return false;
        }
        Observation that = (Observation) o;
        return Objects.equals(id, that.id);
    }


    @Override
    public int hashCode() {
        return Objects.hash(id);
    }
}
