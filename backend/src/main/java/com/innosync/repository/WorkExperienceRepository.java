package com.innosync.repository;

import com.innosync.model.Profile;
import com.innosync.model.WorkExperience;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.transaction.annotation.Transactional;


import java.util.List;

public interface WorkExperienceRepository extends JpaRepository<WorkExperience, Long> {
    List<WorkExperience> findByProfile(Profile profile);


    @Modifying
    @Transactional
    void deleteByProfile(Profile profile);
}
