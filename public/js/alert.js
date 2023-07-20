//skill
document.addEventListener('DOMContentLoaded', function() {
	document.getElementById('skillForm').addEventListener('submit', function(event) {
		event.preventDefault();
		var nameInput = document.getElementById('name');
		var categoryInput = document.getElementById('categoryid');
		var nameError = document.getElementById('nameError');
		var categoryError = document.getElementById('categoryError');

		// Validate Skill Name
		var skillName = nameInput.value.trim();
		if (skillName === '') {
			nameError.textContent = 'Skill Name must be filled';
			nameError.style.display = 'block';
			return; // Stop the form submission
		} else {
			nameError.textContent = '';
			nameError.style.display = 'none';
		}

		// Validate Category
		var categoryId = categoryInput.value.trim();
		if (categoryId === '' || isNaN(categoryId) || categoryId === '0') {
			categoryError.textContent = 'Please select a valid Category';
			categoryError.style.display = 'block';
			return; // Stop the form submission
		} else {
			categoryError.textContent = '';
			categoryError.style.display = 'none';

			// Submit the form
			var form = document.getElementById('skillForm');
			var formData = new FormData(form);

			fetch('/skill/create_skill', {
				method: 'POST',
				body: formData,
			})
				.then(function(response) {
					if (response.ok) {
						// Redirect to the desired page
						window.location.href = '/skill';
					} else {
						return response.json();
					}
				})
				.then(function(data) {
					if (data && data.nameError) {
						nameError.textContent = data.nameError;
						nameError.style.display = 'block';
					}
					if (data && data.categoryError) {
						categoryError.textContent = data.categoryError;
						categoryError.style.display = 'block';
					}
				})
				.catch(function(error) {
					console.log(error);
				});
		}
	});
});
//skill

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

//position
document.addEventListener('DOMContentLoaded', function() {
	document.getElementById('positionForm').addEventListener('submit', function(event) {
		event.preventDefault();
		var nameInput = document.getElementById('name');
		var categoryInput = document.getElementById('categoryid');
		var nameError = document.getElementById('nameError');
		var categoryError = document.getElementById('categoryError');

		// Validate Position Name
		var positionName = nameInput.value.trim();
		if (positionName === '') {
			nameError.textContent = 'Position Name must be filled';
			nameError.style.display = 'block';
			return; // Stop the form submission
		} else {
			nameError.textContent = '';
			nameError.style.display = 'none';
		}

		// Validate Category Position
		var categoryId = categoryInput.value.trim();
		if (categoryId === '' || isNaN(categoryId) || categoryId === '0') {
			categoryError.textContent = 'Please select a valid Position Category';
			categoryError.style.display = 'block';
			return; // Stop the form submission
		} else {
			categoryError.textContent = '';
			categoryError.style.display = 'none';

			// Submit the form
			var form = document.getElementById('positionForm');
			var formData = new FormData(form);

			fetch('/position/create_position', {
				method: 'POST',
				body: formData,
			})
				.then(function(response) {
					if (response.ok) {
						// Redirect to the desired page
						window.location.href = '/position';
					} else {
						return response.json();
					}
				})
				.then(function(data) {
					if (data && data.nameError) {
						nameError.textContent = data.nameError;
						nameError.style.display = 'block';
					}
					if (data && data.categoryError) {
						categoryError.textContent = data.categoryError;
						categoryError.style.display = 'block';
					}
				})
				.catch(function(error) {
					console.log(error);
				});
		}
	});
});
//position

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

//country
document.addEventListener('DOMContentLoaded', function() {
	document.getElementById('countryForm').addEventListener('submit', function(event) {
		event.preventDefault();
		var codeInput = document.getElementById('code');
		var nameInput = document.getElementById('name');
		var codeError = document.getElementById('codeError');
		var nameError = document.getElementById('nameError');

		// Validate Country Code
		var countryCode = codeInput.value.trim();
		if (countryCode === '') {
			codeError.textContent = 'Country Code must be filled';
			codeError.style.display = 'block';
			return; // Stop the form submission

		} else if (countryCode.length > 2) {
			codeError.textContent = 'Country Code must be no more than 2 characters';
			codeError.style.display = 'block';
			return; // Stop the form submission

		} else {
			codeError.textContent = '';
			codeError.style.display = 'none';
		};

		if (nameInput.value.trim() === '') {
			nameError.textContent = 'Country Name must be filled';
			nameError.style.display = 'block';
			return; // Stop the form submission

		} else {
			nameError.textContent = '';
			nameError.style.display = 'none';

			// Submit the form
			var form = document.getElementById('countryForm');
			var formData = new FormData(form);

			fetch('/country/create_country', {
				method: 'POST',
				body: formData,
			})
				.then(function(response) {
					if (response.ok) {
						// Redirect to the desired page
						window.location.href = '/country';
					} else {
						return response.json();
					}
				})
				.then(function(data) {
					if (data && data.codeError) {
						codeError.textContent = data.codeError;
						codeError.style.display = 'block';
					}
					if (data && data.nameError) {
						nameError.textContent = data.nameError;
						nameError.style.display = 'block';
					}
				})
				.catch(function(error) {
					console.log(error);
				});
		}
	});
});
//country

//
