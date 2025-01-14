package posts

import (
	"encoding/json"
	"forum/app/modules"
	"time"
)

var posts = []modules.Post{
	{
		Title:        "Exploring Golang Concurrency",
		Image:        "https://example.com/images/golang-concurrency.jpg",
		Text:         "Concurrency in Go makes it easy to build scalable and efficient software. The unique implementation of goroutines and channels allows developers to manage multiple tasks simultaneously without the complexity often associated with traditional threading models. With goroutines, creating concurrent functions is both lightweight and efficient. Channels further simplify the process by enabling safe communication between concurrent tasks. This approach not only enhances performance but also ensures the code remains clean and maintainable. Whether you're building a web server, data pipeline, or real-time application, Go's concurrency model offers unparalleled advantages.",
		Categories:   []string{"Programming", "Golang", "Concurrency"},
		CreationTime: time.Now().Add(-time.Minute),
		Publisher: modules.User{
			Username:       "srm",
			ProfilePicture: "https://example.com/images/profile/srm.jpg",
			Name:           "Saad",
		},
	},
	{
		Title:        "The Rise of Cloud Computing",
		Image:        "https://example.com/images/cloud-computing.jpg",
		Text:         "Cloud computing has transformed how businesses operate by enabling scalable and on-demand resources. Organizations can now deploy applications globally without investing in physical infrastructure, significantly reducing costs. The flexibility of cloud services, such as Infrastructure as a Service (IaaS) and Platform as a Service (PaaS), allows businesses to scale resources dynamically based on demand. Security and data backup are enhanced with redundant storage options. Additionally, innovations like serverless computing and container orchestration have made cloud adoption even more appealing. The shift to the cloud is more than a trend—it's a cornerstone of digital transformation.",
		Categories:   []string{"Technology", "Cloud", "Innovation"},
		CreationTime: time.Now().Add(-time.Hour),
		Publisher: modules.User{
			Username:       "cloudGuru",
			ProfilePicture: "https://example.com/images/profile/cloudGuru.jpg",
			Name:           "Alice Johnson",
		},
	},
	{
		Title:        "Understanding the Basics of REST APIs",
		Image:        "https://example.com/images/rest-api-basics.jpg",
		Text:         "REST APIs enable seamless communication between applications over the internet. By adhering to principles like statelessness, uniform interfaces, and resource representation, REST simplifies the integration of diverse systems. Developers can perform CRUD operations—Create, Read, Update, and Delete—using standard HTTP methods. For instance, GET retrieves data, POST submits data, PUT updates resources, and DELETE removes them. With JSON or XML as the preferred data formats, REST APIs are both lightweight and developer-friendly. This makes them the backbone of modern web services, allowing for scalable, modular, and interoperable system architectures.",
		Categories:   []string{"Web Development", "APIs", "REST"},
		CreationTime: time.Now().Add(-time.Hour * 24),
		Publisher: modules.User{
			Username:       "apiMaster",
			ProfilePicture: "https://example.com/images/profile/apiMaster.jpg",
			Name:           "Bob Smith",
		},
	},
	{
		Title:        "Artificial Intelligence: Transforming the World",
		Image:        "https://example.com/images/ai-transformation.jpg",
		Text:         "Artificial Intelligence (AI) is reshaping industries by automating processes, improving decision-making, and enhancing user experiences. Machine learning algorithms analyze vast datasets to uncover patterns and make predictions, while natural language processing enables systems to understand and respond to human language. AI is driving innovation in healthcare with predictive diagnostics, in finance with fraud detection, and in entertainment with personalized recommendations. However, ethical considerations around bias, privacy, and accountability remain crucial as AI continues to evolve. Its potential is vast, but so are the responsibilities tied to its deployment.",
		Categories:   []string{"Technology", "AI", "Innovation"},
		CreationTime: time.Now().Add(-time.Hour * 5),
		Publisher: modules.User{
			Username:       "aiExpert",
			ProfilePicture: "https://example.com/images/profile/aiExpert.jpg",
			Name:           "Charlie Brown",
		},
	},
	{
		Title:        "The Importance of Cybersecurity in a Digital World",
		Image:        "https://example.com/images/cybersecurity.jpg",
		Text:         "As businesses and individuals increasingly rely on digital platforms, cybersecurity has become more critical than ever. Cyber threats like phishing, ransomware, and data breaches can compromise sensitive information and disrupt operations. Implementing robust security measures, such as multi-factor authentication, encryption, and regular software updates, is essential. Additionally, fostering a culture of awareness through training helps individuals recognize potential threats. Governments and organizations worldwide are investing in cybersecurity infrastructure to safeguard against sophisticated attacks. In an interconnected world, prioritizing cybersecurity is no longer optional but a necessity to ensure trust and resilience.",
		Categories:   []string{"Technology", "Cybersecurity", "Awareness"},
		CreationTime: time.Now().Add(-time.Hour * 86),
		Publisher: modules.User{
			Username:       "cyberGuard",
			ProfilePicture: "https://example.com/images/profile/cyberGuard.jpg",
			Name:           "Diana Prince",
		},
	},
	{
		Title:        "Healthy Eating: A Key to a Better Life",
		Image:        "https://example.com/images/healthy-eating.jpg",
		Text:         "Healthy eating is fundamental to maintaining physical and mental well-being. A balanced diet that includes fruits, vegetables, lean proteins, and whole grains provides essential nutrients to fuel the body and mind. Proper nutrition can prevent chronic diseases such as diabetes, heart disease, and obesity. Moreover, healthy eating habits enhance energy levels, improve mood, and promote better sleep. While fast food and processed snacks may be convenient, their long-term effects can be detrimental. Planning meals, cooking at home, and staying hydrated are simple yet effective steps to embrace a healthier lifestyle.",
		Categories:   []string{"Health", "Wellness", "Nutrition"},
		CreationTime: time.Now().Add(-time.Hour * 100),
		Publisher: modules.User{
			Username:       "healthGuru",
			ProfilePicture: "https://example.com/images/profile/healthGuru.jpg",
			Name:           "Eva Green",
		},
	},
	{
		Title:        "Space Exploration: Pushing the Boundaries",
		Image:        "https://example.com/images/space-exploration.jpg",
		Text:         "Space exploration continues to push the boundaries of human knowledge and technological capability. From the Moon landings to Mars rovers, each mission expands our understanding of the universe. The International Space Station serves as a hub for scientific research, while advancements in telescopes allow us to study distant galaxies. Private companies like SpaceX and Blue Origin are revolutionizing space travel with reusable rockets and ambitious plans for interplanetary colonization. Despite challenges like funding and technological hurdles, the pursuit of space exploration inspires global collaboration and fuels curiosity about humanity's place in the cosmos.",
		Categories:   []string{"Science", "Space", "Exploration"},
		CreationTime: time.Now().Add(-time.Hour * 12),
		Publisher: modules.User{
			Username:       "spaceExplorer",
			ProfilePicture: "https://example.com/images/profile/spaceExplorer.jpg",
			Name:           "Neil Armstrong",
		},
	},
}

func GetPost(path string) ([]byte, error) {
	return json.Marshal(posts)
}
