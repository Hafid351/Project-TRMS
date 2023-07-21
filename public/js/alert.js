//skill
// document.addEventListener('DOMContentLoaded', function() {
// 	document.getElementById('skillForm').addEventListener('submit', function(event) {
// 		event.preventDefault();
// 		var nameInput = document.getElementById('name');
// 		var categoryInput = document.getElementById('categoryid');
// 		var nameError = document.getElementById('nameError');
// 		var categoryError = document.getElementById('categoryError');

// 		// Validate Skill Name
// 		var skillName = nameInput.value.trim();
// 		if (skillName === '') {
// 			nameError.textContent = 'Skill Name must be filled';
// 			nameError.style.display = 'block';

// 		} else {
// 			nameError.textContent = '';
// 			nameError.style.display = 'none';
// 		}

// 		// Validate Category
// 		var categoryId = categoryInput.value.trim();
// 		if (categoryId === '' || isNaN(categoryId) || categoryId === '0') {
// 			categoryError.textContent = 'Please select a valid Category';
// 			categoryError.style.display = 'block';
			
// 		} else {
// 			categoryError.textContent = '';
// 			categoryError.style.display = 'none';

// 			// Submit the form
// 			var form = document.getElementById('skillForm');
// 			var formData = new FormData(form);

// 			fetch('/skill/create_skill', {
// 				method: 'POST',
// 				body: formData,
// 			})
// 				.then(function(response) {
// 					if (response.ok) {
// 						// Redirect to the desired page
// 						window.location.href = '/skill';
// 					} else {
// 						return response.json();
// 					}
// 				})
// 				.then(function(data) {
// 					if (data && data.nameError) {
// 						nameError.textContent = data.nameError;
// 						nameError.style.display = 'block';
// 					}
// 					if (data && data.categoryError) {
// 						categoryError.textContent = data.categoryError;
// 						categoryError.style.display = 'block';
// 					}
// 				})
// 				.catch(function(error) {
// 					console.log(error);
// 				});
// 		}
// 	});
// });
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
		} else {
			nameError.textContent = '';
			nameError.style.display = 'none';
		}

		// Validate Category
		var categoryId = categoryInput.value.trim();
		if (categoryId === '' || isNaN(categoryId) || categoryId === '0') {
			categoryError.textContent = 'Please select a valid Category';
			categoryError.style.display = 'block';
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
					alert('Skill Created Successfully!'); // Notifikasi pop-up sukses
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

		// Validate Skill Category Name
		var skillCategoryName = nameInput.value.trim();
		if (skillCategoryName === '') {
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
					alert('Skill Category Created Successfully!'); // Notifikasi pop-up sukses
				} else {
					return response.json();
				}
			})
			.then(function(data) {
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
			
		} else {
			nameError.textContent = '';
			nameError.style.display = 'none';
		}

		// Validate Category Position
		var categoryId = categoryInput.value.trim();
		if (categoryId === '' || isNaN(categoryId) || categoryId === '0') {
			categoryError.textContent = 'Please select a valid Position Category';
			categoryError.style.display = 'block';
			
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
						alert('Position Created Successfully!'); // Notifikasi pop-up sukses
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

		// Validate Position Category Name
		var positionCategoryName = nameInput.value.trim();
		if (positionCategoryName === '') {
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
					alert('Position Category Created Successfully!'); // Notifikasi pop-up sukses
				} else {
					return response.json();
				}
			})
			.then(function(data) {
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
//position category

//departement
document.addEventListener('DOMContentLoaded', function() {
	document.getElementById('departementForm').addEventListener('submit', function(event) {
		event.preventDefault();
		var nameInput = document.getElementById('name');
		var nameError = document.getElementById('nameError');

		// Validate Departement Name
		var departementName = nameInput.value.trim();
		if (departementName === '') {
			nameError.textContent = 'Departement Name must be filled';
			nameError.style.display = 'block';
		} else {
			nameError.textContent = '';
			nameError.style.display = 'none';

			// Submit the form
			var form = document.getElementById('departementForm');
			var formData = new FormData(form);

			fetch('/departement/create_departement', {
				method: 'POST',
				body: formData,
			})
			.then(function(response) {
				if (response.ok) {
					// Redirect to the desired page
					window.location.href = '/departement';
					alert('Department Created Successfully!'); // Notifikasi pop-up sukses
				} else {
					return response.json();
				}
			})
			.then(function(data) {
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
			
		} else if (countryCode.length > 2) {
			codeError.textContent = 'Country Code must be no more than 2 characters';
			codeError.style.display = 'block';
			return;

		} else {
			codeError.textContent = '';
			codeError.style.display = 'none';
		};

		if (nameInput.value.trim() === '') {
			nameError.textContent = 'Country Name must be filled';
			nameError.style.display = 'block';

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
						alert('Country Created Successfully!'); // Notifikasi pop-up sukses
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

//language
document.addEventListener('DOMContentLoaded', function() {
	document.getElementById('languageForm').addEventListener('submit', function(event) {
		event.preventDefault();
		var codeInput = document.getElementById('code');
		var nameInput = document.getElementById('name');
		var codeError = document.getElementById('codeError');
		var nameError = document.getElementById('nameError');

		// Validate Language Code
		var languageCode = codeInput.value.trim();
		if (languageCode === '') {
			codeError.textContent = 'Language Code must be filled';
			codeError.style.display = 'block';

		} else {
			codeError.textContent = '';
			codeError.style.display = 'none';
		};

		if (nameInput.value.trim() === '') {
			nameError.textContent = 'Language Name must be filled';
			nameError.style.display = 'block';

		} else {
			nameError.textContent = '';
			nameError.style.display = 'none';

			// Submit the form
			var form = document.getElementById('languageForm');
			var formData = new FormData(form);

			fetch('/language/create_language', {
				method: 'POST',
				body: formData,
			})
				.then(function(response) {
					if (response.ok) {
						// Redirect to the desired page
						window.location.href = '/language';
						alert('Language Created Successfully!'); // Notifikasi pop-up sukses
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
//language

//university
document.addEventListener('DOMContentLoaded', function() {
	document.getElementById('universityForm').addEventListener('submit', function(event) {
		event.preventDefault();
		var nameInput = document.getElementById('name');
		var urlInput = document.getElementById('url');
		var nameError = document.getElementById('nameError');
		var urlError = document.getElementById('urlError');

		// Validate University Name
		var universityName = nameInput.value.trim();
		if (universityName === '') {
			nameError.textContent = 'University Name must be filled';
			nameError.style.display = 'block';

		} else {
			nameError.textContent = '';
			nameError.style.display = 'none';
		};

		// Validate University Web URL
		var universityUrl = urlInput.value.trim();
		if (universityUrl === '') {
			urlError.textContent = 'University Web URL must be filled';
			urlError.style.display = 'block';

		} else {
			urlError.textContent = '';
			urlError.style.display = 'none';

			// Submit the form
			var form = document.getElementById('universityForm');
			var formData = new FormData(form);

			fetch('/university/create_university', {
				method: 'POST',
				body: formData,
			})
				.then(function(response) {
					if (response.ok) {
						// Redirect to the desired page
						window.location.href = '/university';
						alert('University Created Successfully!'); // Notifikasi pop-up sukses
					} else {
						return response.json();
					}
				})
				.then(function(data) {
					if (data && data.nameError) {
						nameError.textContent = data.nameError;
						nameError.style.display = 'block';
					}
					if (data && data.urlError) {
						urlError.textContent = data.urlError;
						urlError.style.display = 'block';
					}
				})
				.catch(function(error) {
					console.log(error);
				});
		}
	});
});

//university

//industry
document.addEventListener('DOMContentLoaded', function() {
	document.getElementById('industryForm').addEventListener('submit', function(event) {
		event.preventDefault();
		var nameInput = document.getElementById('name');
		var nameError = document.getElementById('nameError');

		// Validate Industry Name
		var industryName = nameInput.value.trim();
		if (industryName === '') {
			nameError.textContent = 'Industry Name must be filled';
			nameError.style.display = 'block';

		} else {
			nameError.textContent = '';
			nameError.style.display = 'none';

			// Submit the form
			var form = document.getElementById('industryForm');
			var formData = new FormData(form);

			fetch('/industry/create_industry', {
				method: 'POST',
				body: formData,
			})
				.then(function(response) {
					if (response.ok) {
						// Redirect to the desired page
						window.location.href = '/industry';
						alert('Industry Created Successfully!'); // Notifikasi pop-up sukses
					} else {
						return response.json();
					}
				})
				.then(function(data) {
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
//industry

//company
document.addEventListener('DOMContentLoaded', function() {
	document.getElementById('companyForm').addEventListener('submit', function(event) {
		event.preventDefault();
		var nameInput = document.getElementById('name');
		var industryInput = document.getElementById('industryid');
		var countryInput = document.getElementById('countryid');
		var provinceInput = document.getElementById('provinceid');
		var cityInput = document.getElementById('cityid');
		var nameError = document.getElementById('nameError');
		var industryError = document.getElementById('industryError');
		var countryError = document.getElementById('countryError');
		var provinceError = document.getElementById('provinceError');
		var cityError = document.getElementById('cityError');

		// Validate Company Name
		var companyName = nameInput.value.trim();
		if (companyName === '') {
			nameError.textContent = 'Company Name must be filled';
			nameError.style.display = 'block';

		} else {
			nameError.textContent = '';
			nameError.style.display = 'none';
		}

		// Validate Industry
		var industryId = industryInput.value.trim();
		if (industryId === '' || isNaN(industryId) || industryId === '0') {
			industryError.textContent = 'Please select a valid Industry';
			industryError.style.display = 'block';

		} else {
			industryError.textContent = '';
			industryError.style.display = 'none';
		}

		// Validate Country
		var countryId = countryInput.value.trim();
		if (countryId === '' || isNaN(countryId) || countryId === '0') {
			countryError.textContent = 'Please select a valid Country';
			countryError.style.display = 'block';
			
		} else {
			countryError.textContent = '';
			countryError.style.display = 'none';
		}

		// Validate Province
		var provinceId = provinceInput.value.trim();
		if (provinceId === '' || isNaN(provinceId) || provinceId === '0') {
			provinceError.textContent = 'Please select a valid Province';
			provinceError.style.display = 'block';
			
		} else {
			provinceError.textContent = '';
			provinceError.style.display = 'none';
		}

		// Validate City
		var cityId = cityInput.value.trim();
		if (cityId === '' || isNaN(cityId) || cityId === '0') {
			cityError.textContent = 'Please select a valid City';
			cityError.style.display = 'block';
			
		} else {
			cityError.textContent = '';
			cityError.style.display = 'none';
		}

		// Submit the form
		var form = document.getElementById('companyForm');
		var formData = new FormData(form);

		fetch('/company/create_company', {
			method: 'POST',
			body: formData,
		})
			.then(function(response) {
				if (response.ok) {
					// Redirect to the desired page
					window.location.href = '/company';
					alert('Company Created Successfully!'); // Notifikasi pop-up sukses
				} else {
					return response.json();
				}
			})
			.then(function(data) {
				if (data && data.nameError) {
					nameError.textContent = data.nameError;
					nameError.style.display = 'block';
				}
				// Handle other potential errors from the server if needed
			})
			.catch(function(error) {
				console.log(error);
			});
	});
});

//company

