package com.github.mdeluise.tracky.tracker;

import com.github.mdeluise.tracky.authentication.User;
import com.github.mdeluise.tracky.observation.Observation;
import jakarta.persistence.CascadeType;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.OneToMany;
import jakarta.persistence.Table;
import org.hibernate.validator.constraints.Length;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Objects;

@Entity
@Table(name = "trackers")
public class Tracker {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;
    @Length(max = 30)
    private String name;
    @Length(max = 100)
    private String description;
    @Length(max = 10)
    private String unit;
    @OneToMany(mappedBy = "tracker", cascade = CascadeType.ALL)
    private List<Observation> observations = new ArrayList<>();
    @ManyToOne
    @JoinColumn(name = "user_id", nullable = false)
    private User user;
    private Date lastObservationOn;


    public Long getId() {
        return id;
    }


    public void setId(Long id) {
        this.id = id;
    }


    public String getName() {
        return name;
    }


    public void setName(String name) {
        this.name = name;
    }


    public String getDescription() {
        return description;
    }


    public void setDescription(String description) {
        this.description = description;
    }


    public String getUnit() {
        return unit;
    }


    public void setUnit(String unit) {
        this.unit = unit;
    }


    public List<Observation> getObservations() {
        return observations;
    }


    public void setObservations(List<Observation> observations) {
        this.observations = observations;
    }


    public void addObservation(Observation observation) {
        observations.add(observation);
    }


    public void removeObservation(Observation observation) {
        observations.remove(observation);
    }


    public User getUser() {
        return user;
    }


    public void setUser(User user) {
        this.user = user;
    }


    public Date getLastObservationOn() {
        return lastObservationOn;
    }


    public void setLastObservationOn(Date lastObservationOn) {
        this.lastObservationOn = lastObservationOn;
    }


    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (o == null || getClass() != o.getClass()) {
            return false;
        }
        Tracker tracker = (Tracker) o;
        return Objects.equals(id, tracker.id);
    }


    @Override
    public int hashCode() {
        return Objects.hash(id);
    }
}
