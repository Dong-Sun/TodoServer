function handleSubmit(event) {
	event.preventDefault();
	let toAdd = $('#todoInput').val();
	$('#ul-1').append('<li>' + toAdd + '</li>');
	
	const frm = new FormData();
	frm.append('message',toAdd);

	axios.post('/todo',frm)
		.then(resp => {
		console.log(resp);
	});

}
