"use client";
import React, { useState } from "react";
import styles from "./page.module.css";

export default function ProfileCreation() {
  const [step, setStep] = useState(1);
  const [formData, setFormData] = useState({
    fullName: "",
    email: "",
    telegram: "",
    github: "",
    bio: "",
    position: "",
    technologies: [],
    expertise: "",
    experience: "",
    education: "",
    resume: null,
    workExperiences: [],
  });

  const [workExperience, setWorkExperience] = useState({
    position: "",
    company: "",
    startDate: "",
    endDate: "",
    description: "",
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleWorkChange = (e) => {
    const { name, value } = e.target;
    setWorkExperience((prev) => ({ ...prev, [name]: value }));
  };

  const addTechnology = (tech) => {
    if (!formData.technologies.includes(tech)) {
      setFormData((prev) => ({
        ...prev,
        technologies: [...prev.technologies, tech],
      }));
    }
  };

  const removeTechnology = (tech) => {
    setFormData((prev) => ({
      ...prev,
      technologies: prev.technologies.filter((t) => t !== tech),
    }));
  };

  const addWorkExperience = () => {
    if (
      workExperience.position &&
      workExperience.company &&
      workExperience.startDate &&
      workExperience.endDate &&
      workExperience.description
    ) {
      setFormData((prev) => ({
        ...prev,
        workExperiences: [...prev.workExperiences, workExperience],
      }));
      setWorkExperience({
        position: "",
        company: "",
        startDate: "",
        endDate: "",
        description: "",
      });
    } else {
      alert("Please fill out all fields for work experience.");
    }
  };

  const renderStep1 = () => (
    <div className={styles.profileContainer}>
      <h2 className={styles.stepTitle}>Step 1: Personal Information</h2>
      <div className={styles.inputGroup}>
        <input
          className={styles.inputField}
          name="fullName"
          placeholder="Full Name"
          value={formData.fullName}
          onChange={handleChange}
        />
        <input
          className={styles.inputField}
          name="email"
          placeholder="Email"
          value={formData.email}
          onChange={handleChange}
        />
        <input
          className={styles.inputField}
          name="telegram"
          placeholder="Telegram Handle"
          value={formData.telegram}
          onChange={handleChange}
        />
        <input
          className={styles.inputField}
          name="github"
          placeholder="GitHub Username"
          value={formData.github}
          onChange={handleChange}
        />
        <textarea
          className={`${styles.inputField} ${styles.textArea}`}
          name="bio"
          placeholder="Bio / About Me"
          value={formData.bio}
          onChange={handleChange}
        />
      </div>
      <div className={styles.navButtons}>
        <button type="button" className={styles.navButton} onClick={() => setStep(2)}>
          Next
        </button>
      </div>
    </div>
  );

  const renderStep2 = () => {
    const availableTechnologies = ["React", "Next.js", "Vue", "Figma", "CSS", "Node.js", "Python", "Java"];
    return (
      <div className={styles.profileContainer}>
        <h2 className={styles.stepTitle}>Step 2: Professional Details</h2>
        <div className={styles.inputGroup}>
          <input
            className={styles.inputField}
            name="position"
            placeholder="Current Position"
            value={formData.position}
            onChange={handleChange}
          />

          <div className={styles.techButtons}>
            {availableTechnologies.map((tech) => (
              <button
                key={tech}
                type="button"
                className={
                  formData.technologies.includes(tech)
                    ? `${styles.techButton} ${styles.selected}`
                    : styles.techButton
                }
                onClick={() => {
                  formData.technologies.includes(tech)
                    ? removeTechnology(tech)
                    : addTechnology(tech);
                }}
              >
                {tech}
              </button>
            ))}
          </div>

          <select
            className={styles.select}
            name="expertise"
            value={formData.expertise}
            onChange={handleChange}
          >
            <option value="">Select Expertise Level</option>
            <option value="Beginner">Beginner</option>
            <option value="Intermediate">Intermediate</option>
            <option value="Expert">Expert</option>
          </select>

          <select
            className={styles.select}
            name="experience"
            value={formData.experience}
            onChange={handleChange}
          >
            <option value="">Years of Experience</option>
            <option value="<1">Less than 1 year</option>
            <option value="1-3">1–3 years</option>
            <option value=">3">More than 3 years</option>
          </select>

          <select
            className={styles.select}
            name="education"
            value={formData.education}
            onChange={handleChange}
          >
            <option value="">Highest Degree</option>
            <option value="Bachelor">Bachelor</option>
            <option value="Master">Master</option>
            <option value="PhD">PhD</option>
          </select>
        </div>
        <div className={styles.navButtons}>
          <button type="button" className={styles.navButton} onClick={() => setStep(1)}>
            Back
          </button>
          <button type="button" className={styles.navButton} onClick={() => setStep(3)}>
            Next
          </button>
        </div>
      </div>
    );
  };

  const renderStep3 = () => (
    <div className={styles.profileContainer}>
      <h2 className={styles.stepTitle}>Step 3: Add Work Experience</h2>
      <div className={styles.inputGroup}>
        <input
          className={styles.inputField}
          name="position"
          placeholder="Position"
          value={workExperience.position}
          onChange={handleWorkChange}
        />
        <input
          className={styles.inputField}
          name="company"
          placeholder="Company"
          value={workExperience.company}
          onChange={handleWorkChange}
        />
        <input
          className={styles.inputField}
          name="startDate"
          placeholder="Start Date"
          value={workExperience.startDate}
          onChange={handleWorkChange}
        />
        <input
          className={styles.inputField}
          name="endDate"
          placeholder="End Date"
          value={workExperience.endDate}
          onChange={handleWorkChange}
        />
        <textarea
          className={`${styles.inputField} ${styles.textArea}`}
          name="description"
          placeholder="Work Description"
          value={workExperience.description}
          onChange={handleWorkChange}
        />
      </div>

      <div className={styles.workList}>
        <h4>Added Experiences:</h4>
        {formData.workExperiences.length === 0 && <p>No experiences added yet.</p>}
        <ul>
          {formData.workExperiences.map((exp, index) => (
            <li key={index} className={styles.workItem}>
              <strong>{exp.position}</strong> at {exp.company} ({exp.startDate} – {exp.endDate})
              <p>{exp.description}</p>
            </li>
          ))}
        </ul>
      </div>

      <div className={styles.navButtons}>
        <button type="button" className={styles.navButton} onClick={() => setStep(2)}>
          Back
        </button>
        <button
          type="button"
          className={styles.navButton}
          onClick={() => {
            addWorkExperience();
            alert("Profile Created Successfully!");
          }}
        >
          Finish
        </button>
        <button type="button" className={styles.navButton} onClick={addWorkExperience}>
          Add Another
        </button>
      </div>
    </div>
  );

  return (
    <main className={styles.main}>
      <div className={styles.progressBar}>
        <div className={styles.stepIndicator} style={{ width: step * 33 + "%" }}></div>
      </div>
      {step === 1 && renderStep1()}
      {step === 2 && renderStep2()}
      {step === 3 && renderStep3()}
    </main>
  );
}