function searchAvailabilityByRoom(roomID) {
	document.getElementById("check-availability-button").onclick = () => {
		const html = `
              <form id="check-availability-form" action="" method="POST" novalidate class="needs-validation">
                <div class="form-row">
                  <div class="col">
                    <div class="form-row" id="reservation-dates-modal">
                      <div class="col">
                        <input disabled required class="form-control" type="text" name="start" id="start" placeholder="Arrival">
                      </div>
                      <div class="col">
                        <input disabled required class="form-control" type="text" name="end" id="end" placeholder="Departure">
                      </div>
                    </div>
                  </div>
                </div>
              </form>
            `

		attention.custom({
			msg: html,
			title: "Choose your dates",
			willOpen: () => {
				const datesModal = document.getElementById("reservation-dates-modal")
				const rangePicker = new DateRangePicker(datesModal, {
					showOnFocus: true,
					minDate: new Date(),
				})
			},
			didOpen: () => {
				document.getElementById("start").removeAttribute("disabled")
				document.getElementById("end").removeAttribute("disabled")
			},
			callback: async (result) => {
				const form = document.getElementById("check-availability-form")
				const formData = new FormData(form)
				const tkn = document.querySelector('meta[name="csrf-token"]').content
				formData.append("csrf_token", tkn)
				formData.append("room_id", roomID)

				const response = await fetch("/search-availability-json", {
					method: "POST",
					body: formData,
				})

				const data = await response.json()
				if (data.ok) {
					attention.custom({
						icon: "success",
						msg: `
				      <p>Room is available!</p>
				      <p><a href="/book-room?id=${data.room_id}&s=${data.start_date}&e=${data.end_date}" class="btn btn-primary">Book now!</a></p>`,
						showConfirmButton: false,
					})
				} else {
					attention.error({ msg: "No availability " })
				}
			},
		})
	}
}
