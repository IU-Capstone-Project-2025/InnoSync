## Machine Learning model for team recommendation
simple_model.ipynb consists of 2 parts.
### Mock candidate/job data
First part is mocking candidate and job data for testing and development. I create 500 possible candidates and 30 jobs. The features for the candidate are id, name, main_role, skills, experience (in years), location, current_company, salary_expectation. The features for the job are project_name, description, required_roles, team_composition, required_skills, company, location, budget, and duration. 
### Simple model for recommendation
Model first analyzes job requirements using TF-IDF vectorization to understand skill importance, then applies cosine similarity to match individual candidates to role needs. In future, the model will be changed by a content-based filtering recommendation system.
