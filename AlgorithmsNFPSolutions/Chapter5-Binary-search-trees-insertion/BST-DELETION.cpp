struct node {
	int data;
	node *left, *right;
};

node* delete_node(node *root, int data)
{
	if(root == nullptr) return root;
	else if(data < root->data) root->left = delete_node(root->left, data);
	else if(data > root->data) root->right = delete_node(root->right, data);

	else
	{
		if(root->left == nullptr && root->right == nullptr) // Case 1
		{
			free(root);
			root = nullptr;
		}
		else if(root->left == nullptr) // Case 2
		{
			node* temp = root;
			root = root->right;
			free(temp);
		}
		else if(root->right == nullptr) // Case 2
		{
			node* temp = root;
			root = root->left;
			free(temp);
		}
		else // Case 3 
		{
			node* temp = root->right;

			while(temp->left != nullptr) temp = temp->left;

			root->data = temp->data;
			root->right = delete_node(root->right, temp->data);
		}
	}
	return root;
}
