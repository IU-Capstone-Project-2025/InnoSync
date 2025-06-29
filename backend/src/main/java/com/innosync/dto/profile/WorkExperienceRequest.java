package com.innosync.dto.profile;

import lombok.AllArgsConstructor;
import lombok.Data;

import java.time.LocalDate;

@Data
@AllArgsConstructor
public class WorkExperienceRequest {
    private LocalDate startDate;
    private LocalDate endDate;
    private String position;
    private String company;
    private String description;
}
