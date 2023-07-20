document.addEventListener('DOMContentLoaded', function() {
	document.getElementById('skillCategoryForm').addEventListener('submit', function(event) {
		event.preventDefault();
		var nameInput = document.getElementById('name');
		var nameError = document.getElementById('nameError');

		if (nameInput.value.trim() === '') {
			nameError.textContent = 'Skill Category Name must be filled';
			nameError.style.display = 'block';
		} else {
			nameError.textContent = '';
			nameError.style.display = 'none';

			// Submit the form
			var form = document.getElementById('skillCategoryForm');
			var formData = new FormData(form);

			fetch('/skillcategory/create_skillcategory', {
				method: 'POST',
				body: formData,
			})
				.then(function(response) {
					if (response.ok) {
						// Redirect to the desired page
						window.location.href = '/skillcategory';
					} else {
						return response.json();
					}
				})
				.then(function(data) {
					if (data && data.Error) {
						nameError.textContent = data.Error;
						nameError.style.display = 'block';
					}
				})
				.catch(function(error) {
					console.log(error);
				});
		}
	});
});
