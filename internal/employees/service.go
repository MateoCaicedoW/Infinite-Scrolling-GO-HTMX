package employees

// Mock 50 employees with different roles and skills.
var employeesList = employees{
	{"Alice Johnson", "Software Engineer", "Passionate about backend development.", []string{"Go", "Microservices", "Cloud"}},
	{"Bob Smith", "Frontend Developer", "Loves creating beautiful UI.", []string{"React", "CSS", "TypeScript"}},
	{"Charlie Brown", "DevOps Engineer", "Ensures smooth CI/CD pipelines.", []string{"Docker", "Kubernetes", "Terraform"}},
	{"David Wilson", "Product Manager", "Drives product vision and strategy.", []string{"Agile", "User Research", "Roadmaps"}},
	{"Eve Adams", "Data Scientist", "Expert in AI and ML.", []string{"Python", "TensorFlow", "Data Analysis"}},
	{"Frank White", "Security Engineer", "Focuses on cybersecurity best practices.", []string{"Penetration Testing", "Cryptography", "Compliance"}},
	{"Grace Hall", "UX Designer", "Creates intuitive and accessible designs.", []string{"Figma", "User Experience", "Prototyping"}},
	{"Henry King", "System Administrator", "Manages IT infrastructure.", []string{"Linux", "Networking", "Virtualization"}},
	{"Isabel Scott", "Software Architect", "Designs scalable software solutions.", []string{"Go", "Distributed Systems", "Architecture"}},
	{"Jack Lee", "QA Engineer", "Ensures software quality through testing.", []string{"Selenium", "Test Automation", "CI/CD"}},
	{"Katherine Green", "Marketing Specialist", "Drives digital marketing campaigns.", []string{"SEO", "Content Marketing", "Social Media"}},
	{"Liam Carter", "Technical Writer", "Documents complex systems for users.", []string{"Technical Writing", "Documentation", "API Writing"}},
	{"Mia Lopez", "Business Analyst", "Analyzes business needs and solutions.", []string{"Data Analysis", "Requirements Gathering", "Process Improvement"}},
	{"Noah Allen", "Full Stack Developer", "Builds end-to-end applications.", []string{"Node.js", "React", "MongoDB"}},
	{"Olivia Mitchell", "HR Manager", "Manages employee relations and hiring.", []string{"Recruitment", "Training", "Company Culture"}},
	{"Paul Robinson", "Database Administrator", "Optimizes database performance.", []string{"PostgreSQL", "MySQL", "Indexing"}},
	{"Quinn Harrison", "Mobile Developer", "Creates mobile applications.", []string{"Swift", "Kotlin", "Flutter"}},
	{"Rachel Nelson", "Scrum Master", "Facilitates agile team processes.", []string{"Scrum", "Agile", "Project Management"}},
	{"Samuel Young", "Embedded Systems Engineer", "Develops firmware for embedded devices.", []string{"C", "RTOS", "IoT"}},
	{"Tina Baker", "Network Engineer", "Designs and maintains network infrastructure.", []string{"Cisco", "Firewalls", "VPN"}},
	{"Umar Patel", "Cloud Engineer", "Works with cloud infrastructure.", []string{"AWS", "Azure", "GCP"}},
	{"Victoria Evans", "Blockchain Developer", "Develops blockchain applications.", []string{"Solidity", "Ethereum", "Smart Contracts"}},
	{"William Torres", "Game Developer", "Passionate about game development.", []string{"Unity", "C#", "Game Physics"}},
	{"Xavier Gomez", "AI Engineer", "Builds AI-powered solutions.", []string{"Machine Learning", "Deep Learning", "NLP"}},
	{"Yasmine Rodriguez", "Legal Consultant", "Provides legal guidance in tech.", []string{"Contracts", "Compliance", "IP Law"}},
	{"Zachary Diaz", "Site Reliability Engineer", "Ensures system reliability and uptime.", []string{"SRE", "Observability", "Incident Management"}},
	{"Amber Foster", "Ethical Hacker", "Tests systems for security vulnerabilities.", []string{"Pen Testing", "Bug Bounty", "Cybersecurity"}},
	{"Brandon Hughes", "VR Developer", "Works on virtual reality applications.", []string{"VR", "Unreal Engine", "3D Modeling"}},
	{"Catherine Murphy", "Robotics Engineer", "Develops autonomous robots.", []string{"Robotics", "ROS", "Embedded Systems"}},
	{"Derek Phillips", "Technical Support Engineer", "Provides technical support to customers.", []string{"Help Desk", "Troubleshooting", "Customer Support"}},
	{"Eleanor Griffin", "Biomedical Engineer", "Works on medical device technology.", []string{"Medical Devices", "Biotechnology", "3D Printing"}},
	{"Felix Turner", "Cloud Security Engineer", "Ensures security in cloud environments.", []string{"IAM", "Zero Trust", "Cloud Security"}},
	{"Gabriella White", "E-commerce Specialist", "Manages online sales strategies.", []string{"Shopify", "SEO", "Dropshipping"}},
	{"Harrison Bell", "Game Designer", "Creates engaging game mechanics.", []string{"Game Design", "Level Design", "Storytelling"}},
	{"Ivy Clarke", "AI Researcher", "Explores advancements in artificial intelligence.", []string{"AI Ethics", "Neural Networks", "Reinforcement Learning"}},
	{"Jacob Price", "Penetration Tester", "Tests security vulnerabilities in applications.", []string{"Ethical Hacking", "Red Team", "OWASP"}},
	{"Kara Simmons", "Customer Success Manager", "Enhances user experience and retention.", []string{"Customer Service", "Onboarding", "CRM"}},
	{"Leo Dawson", "Mechanical Engineer", "Designs mechanical systems.", []string{"CAD", "Thermodynamics", "Manufacturing"}},
	{"Madeline Ross", "AR Developer", "Creates augmented reality applications.", []string{"AR", "Unity", "3D Modeling"}},
	{"Nathaniel Brooks", "Biotech Researcher", "Develops biotech solutions.", []string{"CRISPR", "Genomics", "Biopharma"}},
	{"Ophelia Knight", "Cybersecurity Analyst", "Monitors and prevents cyber threats.", []string{"SIEM", "SOC", "Threat Intelligence"}},
	{"Peter Quinn", "SaaS Sales Executive", "Drives software sales.", []string{"B2B Sales", "Lead Generation", "Negotiation"}},
	{"Quincy Barnes", "IoT Engineer", "Develops Internet of Things applications.", []string{"IoT", "Edge Computing", "M2M Communication"}},
	{"Rose Fletcher", "EdTech Specialist", "Improves education through technology.", []string{"E-learning", "Instructional Design", "LMS"}},
	{"Simon Ward", "Automotive Engineer", "Designs modern vehicle systems.", []string{"EV", "Autonomous Vehicles", "ADAS"}},
}

// getEmployees retrieves a paginated list of employees from the employeesList array.
// The offset parameter determines where to start retrieving employees from the array.
// The limit parameter specifies the number of employees to display per page.
func getEmployees(offset, limit int) employees {
	// If the offset is equal to the length of the array, it means there are no more records to fetch.
	if offset >= len(employeesList) {
		return employees{}
	}

	// Calculate the end index for the page.
	// min is the minimum of two integers.
	end := min(offset+limit, len(employeesList))

	return employeesList[offset:end]
}
