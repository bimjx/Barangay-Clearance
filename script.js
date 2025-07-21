// (() => {
//     'use strict';
//     const forms = document.querySelectorAll('form');
//     Array.from(forms).forEach(form => {
//         form.addEventListener('submit', event => {
//             if (!form.checkValidity()) {
//                 // event.preventDefault();
//                 // event.stopPropagation();
//             }
//             form.classList.add('was-validated');
//         }, false);
//     });
// })();


// // Bootstrap validation and modal trigger
// (function () {
//     'use strict'
//     var form = document.querySelector('form');
//     var modal = new bootstrap.Modal(document.getElementById('submittedModal'));
//     form.addEventListener('submit', function (event) {
//         if (!form.checkValidity()) {
//             // event.preventDefault();
//             // event.stopPropagation();
//         } else {
//             // event.preventDefault();
//             modal.show();
//             form.classList.remove('was-validated');
//             form.reset();
//         }
//         form.classList.add('was-validated');
//     }, false);
// })();

// // buttons and forms
// const getStartedBtn = document.getElementById('getStartedBtn');
// const applyNowBtn = document.getElementById('applyNowBtn');
// const loginForm = document.getElementById('loginForm');
// const authModal = document.getElementById('authModal');

// //log in success
// loginForm.addEventListener('submit', function (e) {
//     // e.preventDefault();

//     // hide modal
//     const modal = bootstrap.Modal.getInstance(authModal);
//     modal.hide();

//     // Change hero context
//     getStartedBtn.classList.add('d-none');
//     applyNowBtn.classList.remove('d-non9e');
//     document.getElementById('heroTitle').textContent = 'Welcome!';

//     document.getElementById('heroDesc').textContent = 'You are now logged in. Click Apply Now to start your Barangay Clearance application.';
// });

// // get html
// applyNowBtn.addEventListener('click', function () {
//     window.location.href = '/get';
// });