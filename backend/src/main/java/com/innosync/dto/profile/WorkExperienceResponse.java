package com.innosync.dto.profile;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.LocalDate;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class WorkExperienceResponse {
    private LocalDate startDate;
    private LocalDate endDate;
    private String position;
    private String company;
    private String description;
}
