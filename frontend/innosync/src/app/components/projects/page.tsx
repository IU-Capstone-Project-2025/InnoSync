"use client";
import Navbar from "../navbar/page_logged_in"
import { useState } from "react";
import styles from "./page.module.css";

interface Project {
  id: string;
  title: string;
  company: string;
  location: string;
  salary: string;
  experienceLevel: string;
  type: string;
  description: string;
  logo: string;
  tags: string[];
  daysRemaining: number;
  teamSize: string;
  skills: string[];
}

const ProjectListing = () => {
  const [searchTerm, setSearchTerm] = useState("");
  const [selectedFilters, setSelectedFilters] = useState<string[]>([]);

  // Filter states
  const [experienceLevel, setExperienceLevel] = useState<string[]>([]);
  const [projectType, setProjectType] = useState<string[]>([]);
  const [teamSize, setTeamSize] = useState("");
  const [requiredSkills, setRequiredSkills] = useState<string[]>([]);
  const [newSkill, setNewSkill] = useState("");

  // Mock project data
  const projects: Project[] = [
    {
      id: "1",
      title: "Technical Support Specialist",
      company: "Google",
      location: "Ohio, USA",
      salary: "$35-$45k",
      experienceLevel: "Entry Level",
      type: "Academic",
      description: "4 Days Remaining",
      logo: "🔍",
      tags: ["Full Time"],
      daysRemaining: 4,
      teamSize: "1-2",
      skills: ["React", "UI/UX", "Figma"]
    },
    {
      id: "2",
      title: "UI/UX Designer",
      company: "Adobe",
      location: "Minnesota, USA",
      salary: "$50-$75k",
      experienceLevel: "Junior",
      type: "Business",
      description: "4 Days Remaining",
      logo: "🎨",
      tags: ["Full Time"],
      daysRemaining: 4,
      teamSize: "6-10",
  skills: ["React", "UI/UX", "Figma"]
    },
    {
      id: "3",
      title: "Front End Developer",
      company: "Meta",
      location: "California, USA",
      salary: "$50-$75k",
      experienceLevel: "Mid",
      type: "Business",
      description: "4 Days Remaining",
      logo: "🔧",
      tags: ["Internship"],
      daysRemaining: 4,
      teamSize: "3-5",
  skills: ["React", "UI/UX", "Figma"]
    },
    {
      id: "4",
      title: "Marketing Officer",
      company: "Microsoft",
      location: "Montana, USA",
      salary: "$40-$60k",
      experienceLevel: "Senior",
      type: "Business",
      description: "4 Days Remaining",
      logo: "📈",
      tags: ["Full Time"],
      daysRemaining: 4,
      teamSize: "3-5",
  skills: ["React", "UI/UX", "Figma"]
    },
    {
      id: "5",
      title: "Networking Engineer",
      company: "Cisco",
      location: "Michigan, USA",
      salary: "$90-$100k",
      experienceLevel: "Mid",
      type: "Academic",
      description: "4 Days Remaining",
      logo: "🌐",
      tags: ["Full Time"],
      daysRemaining: 4,
      teamSize: "10+",
  skills: ["React", "UI/UX", "Figma"]
    },
    {
      id: "6",
      title: "Senior UX Designer",
      company: "Apple",
      location: "Forty Lane",
      salary: "2-5 Team Members",
      experienceLevel: "Senior",
      type: "Academic",
      description: "Academic",
      logo: "🍎",
      tags: ["Full Time"],
      daysRemaining: 5,
      teamSize: "1-2",
  skills: ["React", "UI/UX", "Figma"]
    },
    {
      id: "7",
      title: "Junior Graphic Designer",
      company: "Facebook",
      location: "Connecticut, USA",
      salary: "$40-$50k",
      experienceLevel: "Junior",
      type: "Business",
      description: "4 Days Remaining",
      logo: "📘",
      tags: ["Full Time"],
      daysRemaining: 4,
      teamSize: "6-10",
  skills: ["React", "UI/UX", "Figma"]
    },
    {
      id: "8",
      title: "Product Designer",
      company: "Twitter",
      location: "6 Week, Turkey",
      salary: "$50k-$75k",
      experienceLevel: "Entry Level",
      type: "Business",
      description: "4 Days Remaining",
      logo: "🐦",
      tags: ["Full Time"],
      daysRemaining: 4,
      teamSize: "3-5",
  skills: ["React", "UI/UX", "Figma"]
    },
  ];

  const handleFilterToggle = (filter: string) => {
    setSelectedFilters((prev) =>
      prev.includes(filter)
        ? prev.filter((f) => f !== filter)
        : [...prev, filter],
    );
  };

  const handleExperienceLevelChange = (level: string, checked: boolean) => {
    setExperienceLevel((prev) =>
      checked ? [...prev, level] : prev.filter((l) => l !== level),
    );
  };

  const handleProjectTypeChange = (type: string, checked: boolean) => {
    setProjectType((prev) =>
      checked ? [...prev, type] : prev.filter((t) => t !== type),
    );
  };

  const addSkill = () => {
    if (newSkill.trim() && !requiredSkills.includes(newSkill.trim())) {
      setRequiredSkills((prev) => [...prev, newSkill.trim()]);
      setNewSkill("");
    }
  };

  const removeSkill = (skill: string) => {
    setRequiredSkills((prev) => prev.filter((s) => s !== skill));
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === "Enter") {
      addSkill();
    }
  };


  // Filter the projects before displaying
    const filteredProjects = projects.filter((project) => {
    // Experience level
    if (experienceLevel.length > 0 && !experienceLevel.includes(project.experienceLevel)) {
        return false;
    }

    // Project type
    if (projectType.length > 0 && !projectType.includes(project.type)) {
        return false;
    }

    // Team size
    if (teamSize && teamSize !== project.teamSize) {
        return false;
    }

    // Required skills (match if project contains all required skills)
    if (requiredSkills.length > 0) {
        const lowerSkills = project.skills.map(s => s.toLowerCase());
        const hasAllSkills = requiredSkills.every(skill =>
        lowerSkills.includes(skill.toLowerCase())
        );
        if (!hasAllSkills) return false;
    }

    // Search term (in title or company)
    const lowerSearch = searchTerm.toLowerCase();
    if (
        searchTerm &&
        !project.title.toLowerCase().includes(lowerSearch) &&
        !project.company.toLowerCase().includes(lowerSearch)
    ) {
        return false;
    }

    return true;
    });




  return (
    <div className={styles.container}>
        <Navbar/>
        <div className={styles.projectListing}>
            
        {/* Left Sidebar - Filters */}
        <div className={styles.sidebar}>
            <div className={styles.filterSection}>
            <h3 className={styles.filterTitle}>Experience Level</h3>
            <div className={styles.filterOptions}>
                {["Entry Level", "Junior", "Mid", "Senior"].map((level) => (
                <label key={level} className={styles.checkboxLabel}>
                    <input
                    type="checkbox"
                    checked={experienceLevel.includes(level)}
                    onChange={(e) =>
                        handleExperienceLevelChange(level, e.target.checked)
                    }
                    className={styles.checkboxInput}
                    />
                    <span className={styles.checkboxText}>{level}</span>
                </label>
                ))}
            </div>
            </div>

            <div className={styles.filterSection}>
            <h3 className={styles.filterTitle}>Project Type</h3>
            <div className={styles.filterOptions}>
                {["Academic", "Business"].map((type) => (
                <label key={type} className={styles.checkboxLabel}>
                    <input
                    type="checkbox"
                    checked={projectType.includes(type)}
                    onChange={(e) =>
                        handleProjectTypeChange(type, e.target.checked)
                    }
                    className={styles.checkboxInput}
                    />
                    <span className={styles.checkboxText}>{type}</span>
                </label>
                ))}
            </div>
            </div>

            <div className={styles.filterSection}>
            <h3 className={styles.filterTitle}>Team Size</h3>
            <select
                value={teamSize}
                onChange={(e) => setTeamSize(e.target.value)}
                className={styles.teamSizeSelect}
            >
                <option value="">Select team size</option>
                <option value="1-2">1-2 members</option>
                <option value="3-5">3-5 members</option>
                <option value="6-10">6-10 members</option>
                <option value="10+">10+ members</option>
            </select>
            </div>

            <div className={styles.filterSection}>
            <h3 className={styles.filterTitle}>Required Skills</h3>
            <div className={styles.skillsInputContainer}>
                <input
                type="text"
                value={newSkill}
                onChange={(e) => setNewSkill(e.target.value)}
                onKeyPress={handleKeyPress}
                placeholder="Add a skill"
                className={styles.skillInput}
                />
                <button onClick={addSkill} className={styles.addSkillBtn}>
                +
                </button>
            </div>
            <div className={styles.skillsContainer}>
                {requiredSkills.map((skill) => (
                <div key={skill} className={styles.skillTag}>
                    <span>{skill}</span>
                    <button
                    onClick={() => removeSkill(skill)}
                    className={styles.removeSkillBtn}
                    >
                    ×
                    </button>
                </div>
                ))}
            </div>
            </div>
        </div>

        {/* Main Content */}
        <div className={styles.mainContent}>
            {/* Search Bar */}
            <div className={styles.searchContainer}>
            <div className={styles.searchWrapper}>
                <div className={styles.searchIcon}>🔍</div>
                <input
                type="text"
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                placeholder="Search"
                className={styles.searchInput}
                />
            </div>
            </div>

            {/* Filter Tags */}
            <div className={styles.filterTags}>
            {selectedFilters.map((filter) => (
                <div key={filter} className={styles.filterTag}>
                <span>{filter}</span>
                <button
                    onClick={() => handleFilterToggle(filter)}
                    className={styles.removeFilterBtn}
                >
                    ×
                </button>
                </div>
            ))}
            </div>

            {/* Project Cards */}
            <div className={styles.projectsGrid}>
            {filteredProjects.map((project) => (
                <div key={project.id} className={styles.projectCard}>
                <div className={styles.projectHeader}>
                    <div className={styles.companyLogo}>{project.logo}</div>
                    <div className={styles.projectInfo}>
                    <h4 className={styles.projectTitle}>{project.title}</h4>
                    <div className={styles.projectDetails}>
                        <span className={styles.location}>📍 {project.location}</span>
                        <span className={styles.salary}>💰 {project.salary}</span>
                        <span className={styles.experience}>⏰ {project.description}</span>
                    </div>
                    </div>
                    <button className={styles.applyBtn}>Apply Now →</button>
                </div>
                </div>
            ))}
            </div>
        </div>
        </div>
    </div>
  );
};

export default ProjectListing;
