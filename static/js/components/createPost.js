import { checkPost } from "../creatposts.js";
import { back, go } from "../router.js";
import ensureAuth from "../utils/ensureAuth.js";
import { Fetch } from "../utils/fetch.js";
import div from "./native/div.js";
import { input } from "./native/input.js";

export const PostCreationBar = () => {
  const placeholder = "want to share a story?! write here...";
  const createButton = div("create-post");
  createButton.onclick = async () => {
    if (!ensureAuth(true)) {
      return;
    }
    go("/create-post", true);
  };
  return createButton.add(
    input("post-input", placeholder),
    div("create-post-button", "+ Create a Post")
  );
};

export const CreatePost = () => {
  const postCreateView = document.createElement("div");
  postCreateView.className = "postCreateView";
  postCreateView.onclick = (e) => {
    if (e.target === postCreateView) {
      postCreateView.remove();
      back();
    }
  };

  const createLabeledInput = (labelText, inputElement) => {
    const formGroup = document.createElement("div");
    formGroup.className = "formGroup";

    const label = document.createElement("label");
    label.textContent = labelText;
    label.setAttribute("for", inputElement.id);

    formGroup.appendChild(label);
    formGroup.appendChild(inputElement);

    return formGroup;
  };

  const titleInput = document.createElement("input");
  titleInput.className = "titleInput";
  titleInput.type = "text";
  titleInput.placeholder = "Enter title";
  titleInput.id = "titleInput";

  const textInput = document.createElement("textarea");
  textInput.className = "textInput";
  textInput.placeholder = "Enter text";
  textInput.id = "textInput";

  const imageInput = document.createElement("input");
  imageInput.className = "imageInput";
  imageInput.type = "file";
  imageInput.id = "imageInput";

  const categoryDiv = document.createElement("div");
  categoryDiv.className = "categ";

  const categories = ["Technology", "Sport", "Finance", "Science"];
  categories.forEach((cat) => {
    const checkboxLabel = document.createElement("label");
    checkboxLabel.className = "category-label";

    const checkbox = document.createElement("input");
    checkbox.type = "checkbox";
    checkbox.value = cat;
    checkbox.className = "category-checkbox";

    checkboxLabel.appendChild(checkbox);
    checkboxLabel.appendChild(document.createTextNode(cat));
    categoryDiv.appendChild(checkboxLabel);
  });

  const submitButton = document.createElement("button");
  submitButton.className = "submitButton";
  submitButton.textContent = "Create a Post";
  submitButton.onclick = () => {
    Fetch("/api/posts", {
      method: "POST",
      body: JSON.stringify({
        title: titleInput.value,
        content: textInput.value,
      }),
    });
  };

  const cancelButton = document.createElement("button");
  cancelButton.className = "cancelButton";
  cancelButton.textContent = "Cancel";
  cancelButton.onclick = () => {
    postCreateView.remove();
    back();
  };

  const buttonContainer = document.createElement("div");
  buttonContainer.className = "buttonContainer";
  buttonContainer.appendChild(cancelButton);
  buttonContainer.appendChild(submitButton);

  const postForm = document.createElement("div");
  postForm.className = "postForm";
  postForm.appendChild(createLabeledInput("Title", titleInput));
  postForm.appendChild(createLabeledInput("Description", textInput));
  postForm.appendChild(createLabeledInput("Upload Image", imageInput));
  postForm.appendChild(categoryDiv);
  postForm.appendChild(buttonContainer);

  postCreateView.appendChild(postForm);
  checkPost();
  return postCreateView;
};

export default CreatePost;
