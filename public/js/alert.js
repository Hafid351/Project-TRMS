//skill category
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
						return response.json(); // Parse the response data
					} else {
						return response.json(); // Parse the response data
					}
				})
				.then(function(data) {
					if (data && data.Error) {
						nameError.textContent = data.Error;
						nameError.style.display = 'block';
						alert(data.Error);
					} else {
						// If there is no error, display success message
						alert('Skill Category created successfully');
					}
				})
				.catch(function(error) {
					console.log(error);
					alert('Failed to create Skill Category');
				});
		}
	});
});
//skill category

//position category
document.addEventListener('DOMContentLoaded', function() {
	document.getElementById('positionCategoryForm').addEventListener('submit', function(event) {
		event.preventDefault();
		var nameInput = document.getElementById('name');
		var nameError = document.getElementById('nameError');

		if (nameInput.value.trim() === '') {
			nameError.textContent = 'Position Category Name must be filled';
			nameError.style.display = 'block';
		} else {
			nameError.textContent = '';
			nameError.style.display = 'none';

			// Submit the form
			var form = document.getElementById('positionCategoryForm');
			var formData = new FormData(form);

			fetch('/positioncategory/create_positioncategory', {
				method: 'POST',
				body: formData,
			})
				.then(function(response) {
					if (response.ok) {
						// Redirect to the desired page
						window.location.href = '/positioncategory';
						return response.json(); // Parse the response data
					} else {
						return response.json(); // Parse the response data
					}
				})
				.then(function(data) {
					if (data && data.Error) {
						nameError.textContent = data.Error;
						nameError.style.display = 'block';
						alert(data.Error);
					} else {
						// If there is no error, display success message
						alert('Position Category created successfully');
					}
				})
				.catch(function(error) {
					console.log(error);
					alert('Failed to create Skill Category');
				});
		}
	});
});
//position category

//departement mohon bantuannya
// document.addEventListener('DOMContentLoaded', function() {
// 	document.getElementById('departementForm').addEventListener('submit', function(event) {
// 		event.preventDefault();
// 		var nameInput = document.getElementById('name');
// 		var nameError = document.getElementById('nameError');

// 		if (nameInput.value.trim() === '') {
// 			nameError.textContent = 'Departement Name must be filled';
// 			nameError.style.display = 'block';
// 		} else {
// 			nameError.textContent = '';
// 			nameError.style.display = 'none';

// 			// Submit the form
// 			var form = document.getElementById('departementForm');
// 			var formData = new FormData(form);

// 			fetch('/departement/create_departement', {
// 				method: 'POST',
// 				body: formData,
// 			})
// 				.then(function(response) {
// 					if (response.ok) {
// 						// Redirect to the desired page
// 						window.location.href = '/departement';
// 						return response.json(); // Parse the response data
// 					} else {
// 						return response.json(); // Parse the response data
// 					}
// 				})
// 				.then(function(data) {
// 					if (data && data.Error) {
// 						nameError.textContent = data.Error;
// 						nameError.style.display = 'block';
// 						alert(data.Error);
// 					} else {
// 						// If there is no error, display success message
// 						alert('Departement created successfully');
// 					}
// 				})
// 				.catch(function(error) {
// 					console.log(error);
// 					alert('Failed to create Departement');
// 				});
// 		}
// 	});
// });
//departement