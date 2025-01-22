package posts

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/app/modules"
	"log"
	"strconv"
	"strings"
)

// var posts = []modules.Post{
// 	{
// 		Title:        "Exploring Golang Concurrency",
// 		Image:        "https://example.com/images/golang-concurrency.jpg",
// 		Text:         "Concurrency in Go makes it easy to build scalable and efficient software. The unique implementation of goroutines and channels allows developers to manage multiple tasks simultaneously without the complexity often associated with traditional threading models. With goroutines, creating concurrent functions is both lightweight and efficient. Channels further simplify the process by enabling safe communication between concurrent tasks. This approach not only enhances performance but also ensures the code remains clean and maintainable. Whether you're building a web server, data pipeline, or real-time application, Go's concurrency model offers unparalleled advantages.",
// 		ID:           "posts-3",
// 		Categories:   []string{"Programming", "Golang", "Concurrency"},
// 		CreationTime: time.Now().Add(-time.Minute),
// 		Publisher: modules.User{
// 			Username: "srm",
// 		},
// 	},
// 	{
// 		Title:        "The Rise of Cloud Computing",
// 		Image:        "https://example.com/images/cloud-computing.jpg",
// 		Text:         "Cloud computing has transformed how businesses operate by enabling scalable and on-demand resources. Organizations can now deploy applications globally without investing in physical infrastructure, significantly reducing costs. The flexibility of cloud services, such as Infrastructure as a Service (IaaS) and Platform as a Service (PaaS), allows businesses to scale resources dynamically based on demand. Security and data backup are enhanced with redundant storage options. Additionally, innovations like serverless computing and container orchestration have made cloud adoption even more appealing. The shift to the cloud is more than a trend—it's a cornerstone of digital transformation.",
// 		Categories:   []string{"Technology", "Cloud", "Innovation"},
// 		CreationTime: time.Now().Add(-time.Hour),
// 		Publisher: modules.User{
// 			Username: "cloudGuru",
// 		},
// 	},
// 	{
// 		Title:        "Understanding the Basics of REST APIs",
// 		Image:        "https://example.com/images/rest-api-basics.jpg",
// 		Text:         "REST APIs enable seamless communication between applications over the internet. By adhering to principles like statelessness, uniform interfaces, and resource representation, REST simplifies the integration of diverse systems. Developers can perform CRUD operations—Create, Read, Update, and Delete—using standard HTTP methods. For instance, GET retrieves data, POST submits data, PUT updates resources, and DELETE removes them. With JSON or XML as the preferred data formats, REST APIs are both lightweight and developer-friendly. This makes them the backbone of modern web services, allowing for scalable, modular, and interoperable system architectures.",
// 		Categories:   []string{"Web Development", "APIs", "REST"},
// 		CreationTime: time.Now().Add(-time.Hour * 24),
// 		Publisher: modules.User{
// 			Username: "apiMaster",
// 		},
// 	},
// 	{
// 		Title:        "Artificial Intelligence: Transforming the World",
// 		Image:        "https://example.com/images/ai-transformation.jpg",
// 		Text:         "Artificial Intelligence (AI) is reshaping industries by automating processes, improving decision-making, and enhancing user experiences. Machine learning algorithms analyze vast datasets to uncover patterns and make predictions, while natural language processing enables systems to understand and respond to human language. AI is driving innovation in healthcare with predictive diagnostics, in finance with fraud detection, and in entertainment with personalized recommendations. However, ethical considerations around bias, privacy, and accountability remain crucial as AI continues to evolve. Its potential is vast, but so are the responsibilities tied to its deployment.",
// 		Categories:   []string{"Technology", "AI", "Innovation"},
// 		CreationTime: time.Now().Add(-time.Hour * 5),
// 		Publisher: modules.User{
// 			Username: "aiExpert",
// 		},
// 	},
// 	{
// 		Title:        "The Importance of Cybersecurity in a Digital World",
// 		Image:        "https://example.com/images/cybersecurity.jpg",
// 		Text:         "As businesses and individuals increasingly rely on digital platforms, cybersecurity has become more critical than ever. Cyber threats like phishing, ransomware, and data breaches can compromise sensitive information and disrupt operations. Implementing robust security measures, such as multi-factor authentication, encryption, and regular software updates, is essential. Additionally, fostering a culture of awareness through training helps individuals recognize potential threats. Governments and organizations worldwide are investing in cybersecurity infrastructure to safeguard against sophisticated attacks. In an interconnected world, prioritizing cybersecurity is no longer optional but a necessity to ensure trust and resilience.",
// 		Categories:   []string{"Technology", "Cybersecurity", "Awareness"},
// 		CreationTime: time.Now().Add(-time.Hour * 86),
// 		Publisher: modules.User{
// 			Username: "cyberGuard",
// 		},
// 	},
// 	{
// 		Title:        "Healthy Eating: A Key to a Better Life",
// 		Image:        "https://example.com/images/healthy-eating.jpg",
// 		Text:         "Healthy eating is fundamental to maintaining physical and mental well-being. A balanced diet that includes fruits, vegetables, lean proteins, and whole grains provides essential nutrients to fuel the body and mind. Proper nutrition can prevent chronic diseases such as diabetes, heart disease, and obesity. Moreover, healthy eating habits enhance energy levels, improve mood, and promote better sleep. While fast food and processed snacks may be convenient, their long-term effects can be detrimental. Planning meals, cooking at home, and staying hydrated are simple yet effective steps to embrace a healthier lifestyle.",
// 		Categories:   []string{"Health", "Wellness", "Nutrition"},
// 		CreationTime: time.Now().Add(-time.Hour * 100),
// 		Publisher: modules.User{
// 			Username: "healthGuru",
// 		},
// 	},
// 	{
// 		Title:        "Space Exploration: Pushing the Boundaries",
// 		Image:        "https://example.com/images/space-exploration.jpg",
// 		Text:         "Space exploration continues to push the boundaries of human knowledge and technological capability. From the Moon landings to Mars rovers, each mission expands our understanding of the universe. The International Space Station serves as a hub for scientific research, while advancements in telescopes allow us to study distant galaxies. Private companies like SpaceX and Blue Origin are revolutionizing space travel with reusable rockets and ambitious plans for interplanetary colonization. Despite challenges like funding and technological hurdles, the pursuit of space exploration inspires global collaboration and fuels curiosity about humanity's place in the cosmos.",
// 		Categories:   []string{"Science", "Space", "Exploration"},
// 		CreationTime: time.Now().Add(-time.Hour * 12),
// 		Publisher: modules.User{
// 			Username: "spaceExplorer",
// 		},
// 	},
// }

var posts = []modules.Post{
	// your predefined posts
}

func GetPosts(conn *modules.Connection, forumDB *sql.DB) ([]byte, error) {
	var posts []modules.Post
	var categories []string
	var err error
	var pageStr = ""
	log.Printf("Request URL Path: %s", conn.Req.URL.Path)

	if len(strings.Split(conn.Req.URL.Path, "/")) == 4 {
		pageStr = strings.Split(conn.Req.URL.Path, "/")[3]
		log.Printf("Page: %s", pageStr)
	}
	page := 1

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if strings.Contains(conn.Req.URL.Path, "categories") {
		cat := strings.Split(conn.Req.URL.Path, "/categories=")[1]
		categories = strings.Split(cat, "&")
		log.Printf("Categories found in URL: %v", categories)
	}

	if len(categories) > 0 {
		log.Printf("Fetching posts by categories: %v for page %d", categories, page)
		err = fetchPostsByCategories(categories, &posts, page, forumDB)
	} else {
		log.Printf("Fetching posts for page %d", page)
		err = fetchAllPosts(&posts, page, forumDB)
	}

	if err != nil {
		log.Printf("Error fetching posts: %v", err)
		return nil, fmt.Errorf("error fetching posts: %v", err)
	}

	postJSON, err := json.Marshal(posts)
	if err != nil {
		log.Printf("Error marshaling posts: %v", err)
		return nil, fmt.Errorf("error marshaling posts: %v", err)
	}

	log.Printf("Number of posts returned: %d", len(posts))

	return postJSON, nil
}

func fetchPostsByCategories(categories []string, posts *[]modules.Post, page int, forumDB *sql.DB) error {
	const limit = 10
	offset := (page - 1) * limit

	categoriesCode := GetCategoryMask(categories)
	categoryInt, err := strconv.ParseInt(categoriesCode, 2, 64)
	if err != nil {
		return fmt.Errorf("error parsing category mask: %v", err)
	}

	query := `SELECT user_id, item_id, title, content, categories, created_at 
              FROM posts 
              WHERE (CAST(categories AS INTEGER) & ?) != 0 
              ORDER BY created_at DESC 
              LIMIT ? OFFSET ?`

	log.Printf("Executing query: %s with category mask: %d, limit: %d, offset: %d", query, categoryInt, limit, offset)

	// db, err := sql.Open("sqlite3", "./forum.db")
	// if err != nil {
	// 	log.Printf("Error opening database: %v", err)
	// 	return err
	// }
	// defer db.Close()

	rows, err := forumDB.Query(query, categoryInt, limit, offset)
	if err != nil {
		log.Printf("Error querying posts by category: %v", err)
		return fmt.Errorf("error querying posts by category: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post modules.Post
		var categoryMask string

		if err := rows.Scan(&post.Publisher.Username, &post.ID, &post.Title, &post.Text, &categoryMask, &post.CreationTime); err != nil {
			log.Printf("Error scanning post: %v", err)
			return fmt.Errorf("error scanning post: %v", err)
		}

		post.Categories = GetCategoriesFromMask(categoryMask)
		*posts = append(*posts, post)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return err
	}

	return nil
}

func fetchAllPosts(posts *[]modules.Post, page int, forumDB *sql.DB) error {
	const limit = 10
	offset := (page - 1) * limit

	query := `SELECT user_id, item_id, title, content, categories, created_at 
              FROM posts 
              ORDER BY created_at DESC 
              LIMIT ? OFFSET ?`

	log.Printf("Executing query: %s with limit %d and offset %d", query, limit, offset)

	// db, err := sql.Open("sqlite3", "./forum.db")
	// if err != nil {
	// 	log.Printf("Error opening database: %v", err)
	// 	return err
	// }
	// defer db.Close()

	rows, err := forumDB.Query(query, limit, offset)
	if err != nil {
		log.Printf("Error querying all posts: %v", err)
		return fmt.Errorf("error querying all posts: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post modules.Post
		var categoryMask string

		if err := rows.Scan(&post.Publisher.Username, &post.ID, &post.Title, &post.Text, &categoryMask, &post.CreationTime); err != nil {
			log.Printf("Error scanning post: %v", err)
			return fmt.Errorf("error scanning post: %v", err)
		}

		post.Categories = GetCategoriesFromMask(categoryMask)

		*posts = append(*posts, post)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return err
	}

	return nil
}

func GetCategoriesFromMask(mask string) []string {
	categoryMap := map[string]int{
		"Sport":      0,
		"Technology": 1,
		"Finance":    2,
		"Science":    3,
	}

	categories := []string{}
	for i, c := range mask {
		if c == '1' {
			for category, idx := range categoryMap {
				if idx == i {
					categories = append(categories, category)
					break
				}
			}
		}
	}

	return categories
}
