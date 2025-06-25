package com.innosync.model;

import jakarta.persistence.*;
import lombok.*;

import java.util.List;
import java.util.ArrayList;


@Entity
@Table(name = "user_profile")
@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class Profile {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @OneToOne
    @JoinColumn(name = "user_id", nullable = false, unique = true)
    private User user;

    private String telegram;

    private String github;

    @Column(length = 1000)
    private String bio;

    private String position;

    @Enumerated(EnumType.STRING)
    @Column(nullable = false)
    private Education education;

    private String expertise;

    @Enumerated(EnumType.STRING)
    @Column(name = "expertise_level", nullable = false)
    private ExpertiseLevel expertiseLevel;

    private String resume; // URL or path

    @ManyToMany
    @JoinTable(
            name = "user_profile_technology",
            joinColumns = @JoinColumn(name = "user_profile_id"),
            inverseJoinColumns = @JoinColumn(name = "technology_id")
    )
    private List<Technology> technologies = new ArrayList<>();

}
