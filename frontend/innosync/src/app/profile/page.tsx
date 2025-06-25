"use client";
import React, { useState } from "react";
import styles from "./page.module.css";

function ProfileCreation() {
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

  const addWorkExperience = () => {
    setFormData((prev) => ({
      ...prev,
      workExperiences: [...prev.workExperiences, workExperience],
    }));
    setWorkExperience({ position: "", company: "", startDate: "", endDate: "", description: "" });
  };

  const renderStep1 = () => (
    <div className={styles.profileContainer}>
      <h2 className={styles.stepTitle}>Step 1: Profile Info</h2>
      <div className={styles.inputGroup}>
        <input className={styles.inputField} name="fullName" placeholder="Full Name" value={formData.fullName} onChange={handleChange} />
        <input className={styles.inputField} name="email" placeholder="Email" value={formData.email} onChange={handleChange} />
        <input className={styles.inputField} name="telegram" placeholder="Telegram" value={formData.telegram} onChange={handleChange} />
        <input className={styles.inputField} name="github" placeholder="GitHub" value={formData.github} onChange={handleChange} />
        <textarea className={`${styles.inputField} ${styles.textArea}`} name="bio" placeholder="Bio" value={formData.bio} onChange={handleChange} />
      </div>
      <div className={styles.navButtons}>
        <button className={styles.navButton} onClick={() => setStep(2)}>Next</button>
      </div>
    </div>
  );

  const renderStep2 = () => (
    <div className={styles.profileContainer}>
      <h2 className={styles.stepTitle}>Step 2: Professional Info</h2>
      <div className={styles.inputGroup}>
        <input className={styles.inputField} name="position" placeholder="Position" value={formData.position} onChange={handleChange} />

        <div className={styles.techButtons}>
          {["React", "Next.js", "Vue", "Figma", "CSS"].map((tech) => (
            <button key={tech} className={styles.techButton} onClick={() => addTechnology(tech)}>{tech}</button>
          ))}
        </div>
        <div className={styles.selectedTechs}>{formData.technologies.join(", ")}</div>

        <select className={styles.select} name="expertise" value={formData.expertise} onChange={handleChange}>
          <option value="">Expertise</option>
          <option value="Beginner">Beginner</option>
          <option value="Intermediate">Intermediate</option>
          <option value="Expert">Expert</option>
        </select>

        <select className={styles.select} name="experience" value={formData.experience} onChange={handleChange}>
          <option value="">Years</option>
          <option value="<1">Less than 1</option>
          <option value="1-3">1-3</option>
          <option value=">3">More than 3</option>
        </select>

        <select className={styles.select} name="education" value={formData.education} onChange={handleChange}>
          <option value="">Degree</option>
          <option value="Bachelor">Bachelor</option>
          <option value="Master">Master</option>
          <option value="PhD">PhD</option>
        </select>
      </div>

      <div className={styles.navButtons}>
        <button className={styles.navButton} onClick={() => setStep(1)}>Back</button>
        <button className={styles.navButton} onClick={() => setStep(3)}>Next</button>
      </div>
    </div>
  );

  const renderStep3 = () => (
    <div className={styles.profileContainer}>
      <h2 className={styles.stepTitle}>Step 3: Work Experience</h2>
      <div className={styles.inputGroup}>
        <input className={styles.inputField} name="position" placeholder="Position" value={workExperience.position} onChange={handleWorkChange} />
        <input className={styles.inputField} name="company" placeholder="Company" value={workExperience.company} onChange={handleWorkChange} />
        <input className={styles.inputField} name="startDate" placeholder="Start Date" value={workExperience.startDate} onChange={handleWorkChange} />
        <input className={styles.inputField} name="endDate" placeholder="End Date" value={workExperience.endDate} onChange={handleWorkChange} />
        <textarea className={`${styles.inputField} ${styles.textArea}`} name="description" placeholder="Work Description" value={workExperience.description} onChange={handleWorkChange} />
      </div>

      <div className={styles.navButtons}>
        <button className={styles.navButton} onClick={() => setStep(2)}>Back</button>
        <button className={styles.navButton} onClick={() => { addWorkExperience(); alert("Profile Created Successfully!"); }}>Finish</button>
        <button className={styles.navButton} onClick={addWorkExperience}>Add Another</button>
      </div>
    </div>
  );

  return (
    <div>
      {step === 1 && renderStep1()}
      {step === 2 && renderStep2()}
      {step === 3 && renderStep3()}
    </div>
  );
}

export default ProfileCreation;
