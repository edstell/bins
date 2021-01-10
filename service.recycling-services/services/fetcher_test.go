package services

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	recyclingservicesproto "github.com/edstell/lambda/service.recycling-services/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestWebScraper(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.Write(html)
	}))
	defer server.Close()
	scraper := WebScraper(server.Client(), ParseHTML, server.URL)
	result, err := scraper.Fetch(context.Background(), "property_id")
	require.NoError(t, err)
	fmt.Println(result)
	for i, service := range []*recyclingservicesproto.Service{
		{
			Name:        "Non-recyclable refuse",
			Status:      "Not completed.",
			Schedule:    "Thursday every other week",
			LastService: timestamppb.New(time.Date(2020, 12, 17, 0, 0, 0, 0, time.UTC)),
			NextService: timestamppb.New(time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)),
		},
		{
			Name:        "Paper and cardboard",
			Status:      "Not completed.",
			Schedule:    "Thursday every other week",
			LastService: timestamppb.New(time.Date(2020, 12, 17, 0, 0, 0, 0, time.UTC)),
			NextService: timestamppb.New(time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)),
		},
		{
			Name:        "Green Garden Waste (Subscription)",
			Status:      "Your road was completed on 09/12/2020 at 09:27.",
			Schedule:    "Wednesday every 4th week",
			LastService: timestamppb.New(time.Date(2020, 12, 9, 0, 0, 0, 0, time.UTC)),
			NextService: timestamppb.New(time.Date(2021, 1, 20, 0, 0, 0, 0, time.UTC)),
		},
		{
			Name:        "Food waste",
			Status:      "Not completed.",
			Schedule:    "Thursday every week",
			LastService: timestamppb.New(time.Date(2020, 12, 24, 0, 0, 0, 0, time.UTC)),
			NextService: timestamppb.New(time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)),
		},
		{
			Name:        "Plastic, glass and tins",
			Status:      "Your road was completed on 24/12/2020 at 08:49.",
			Schedule:    "Thursday every other week",
			LastService: timestamppb.New(time.Date(2020, 12, 24, 0, 0, 0, 0, time.UTC)),
			NextService: timestamppb.New(time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC)),
		},
		{
			Name: "Batteries, small electrical items and textiles",
		},
	} {
		assert.Equal(t, service, result[i], service.Name)
	}
}

var html = []byte(`<!DOCTYPE html>
<!--[if lt IE 7]>
<html lang="en" class="no-js lt-ie9 lt-ie8 lt-ie7"> <![endif]-->
<!--[if IE 7]>
<html lang="en" class="no-js lt-ie9 lt-ie8"> <![endif]-->
<!--[if IE 8]>
<html lang="en" class="no-js lt-ie9"> <![endif]-->
<!--[if gt IE 8]><!-->
<html lang="en" class="no-js">
<!--<![endif]-->
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
	<title>Property Results | Resident Service Portal</title>
	<meta name="description" content="">
	<meta name="viewport" content="width=device-width, initial-scale=1">
			<!-- Global site tag (gtag.js) - Google Analytics -->
		<script async src="https://www.googletagmanager.com/gtag/js?id=UA-32549068-3"></script>
		<script>
			window.dataLayer = window.dataLayer || [];
			function gtag() {
				dataLayer.push(arguments);
			}
			gtag('js', new Date());

			gtag('config', 'UA-32549068-3');
		</script>
		
	
	<link rel="stylesheet" href="/system/css/bootstrap.min.css">
	<link rel="stylesheet" href="/system/css/bootstrap-theme.min.css">
	<link rel="stylesheet" href="/system/css/bootstrap-datetimepicker.css">
	<link rel="stylesheet" href="/system/css/main.css">
		<style>
			@keyframes pulse {
				0% {
					opacity: 1;
					content: "Loading";
				}

				10% {
					opacity: 0.5;
				}

				20% {
					opacity: 1;
					content: "» Loading «";
				}

				30% {
					opacity: 0.5;
				}

				40% {
					opacity: 1;
					content: "»» Loading ««";
				}

				50% {
					opacity: 0.5;
				}

				60% {
					opacity: 1;
					content: "»»» Loading «««";
				}

				70% {
					opacity: 0.5;
				}

				80% {
					opacity: 1;
					content: "»»»» Loading ««««";
				}

				90% {
					opacity: 0.5;
				}

				100% {
					opacity: 1;
					content: "»»»»» Loading «««««";
				}
			}
		</style>
	<link rel="stylesheet" href="/cl/css/bromley.css">		<!-- HTML5 shim and Respond.js IE8 support of HTML5 elements and media queries -->
		<!--[if lt IE 9]>
		<script src="//html5shiv.googlecode.com/svn/trunk/html5.js"></script>
		<script src="//oss.maxcdn.com/libs/respond.js/1.3.0/respond.min.js"></script>
		<![endif]-->

		<script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
		<script src="//code.jquery.com/ui/1.10.4/jquery-ui.js"></script>

		<script src="/system/js/vendor/bootstrap.min.js"></script>
				<script src="/system/js/main.js"></script>
	  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css" media="all"></head><body class="page-property " onunload="">
<a href="#main" class="direct bypass-block-link visually-hidden visible-when-focused">Skip Navigation</a>
<!--[if lt IE 7]>
<p class="chromeframe">You are using an <strong>outdated</strong> browser. Please <a href="http://browsehappy.com/">upgrade your browser</a> or <a href="http://www.google.com/chromeframe/?redirect=true">activate Google Chrome Frame</a> to improve your experience.</p>
<![endif]-->
<header class="navbar navbar-default navbar-static-top">
  <div class="container">
    <div class="navbar-header">
      <ul>
      <li class="logo-wrapper"><a href="/"><img src="/cl/img/logo.png" alt="Veolia - In partnership with Bromley" id="logo"></a></li>      </ul>
    </div>
  </div>
  </header>
<div class="page-heading">
	<div class="container">
				<h1 style="display: inline-block;">Property Results</h1>
			</div>
</div><div id="main">
  <div class="container message-area-wrapper"><div class="message-area"></div></div><div class="container" style="position: relative;">
   <div class="row">
      <div class="col-sm-3">
         <div class="left-panel property-address">
            <h3>Address</h3>
            <address><strong>18 Mill Vale, Bromley, BR2 0EW</strong></address>
            <h3>Links</h3>
            <ul>
               <li><a href="https://www.bromley.gov.uk/info/200084/recycling_rubbish_and_waste">Home</a></li>
               <li><a href="/property">Change address</a></li>
            </ul>
         </div>
                    </div>
      <div class="col-sm-9">
                           <div class="results-table-wrapper">
            		<div class="service-wrapper service-id-531 task-id-3212">
		<h3 class="service-name">Non-recyclable refuse					</h3>
		<div class="service-content">
			<div class="image-wrapper">
			<div class="image">

			</div>
			</div>
			<table class="table">
				
				<thead>
				<tr>
					<th class="schedule">Schedule</th>
					<th class="last-service">Last Service</th>
					<th class="next-service">Next Service</th>
				</tr>
				</thead>
				<tbody>
				<tr>
				
					<td class="schedule">
						<span class="table-label">Schedule</span>
						<div></div><div>Thursday every other week</div>					</td>

					<td class="last-service">
						<span class="table-label">Last Service</span>
						17/12/2020					</td>

					<td class="next-service">
						<span class="table-label">Next Service</span>
						04/01/2021					</td>
				</tr>
				<tr>
					<td colspan="3">
													
						<div class="task-state state-not-completed">
							<span class="indicator"></span>
							<p><strong>Last collection: Not completed.</strong></p>						</div>
					</td>
				</tr>
				<tr>
					<td colspan="3" class="event-actions">
						<div class="other-wrapper"><div class="btn-wrapper"><a class="btn btn-link" href="#">Feedback</a></div><ul class="other-dropdown"><li"><a class="btn btn-link" href="/praise/2162/property/100020406685/service/531">Crew Behaviour</a></li><li"><a class="btn btn-link" href="/praise/2159/property/100020406685/service/531">Damage to 3rd Party Vehicle</a></li><li"><a class="btn btn-link" href="/praise/2163/property/100020406685/service/531">Damage to Property</a></li><li"><a class="btn btn-link" href="/praise/2148/property/100020406685/service/531">General Enquiry</a></li><li"><a class="btn btn-link" href="/praise/2186/property/100020406685/service/531">Wrongful Removal</a></li></ul></div>					</td>
				</tr>
								<tr>
					<td colspan="3">
						Please note that missed collections can only be reported within 2 working days of your scheduled collection day.
					</td>
				</tr>
								</tbody>
			</table>
		</div>
	</div>
		<div class="service-wrapper service-id-537 task-id-3218">
		<h3 class="service-name">Paper and cardboard					</h3>
		<div class="service-content">
			<div class="image-wrapper">
			<div class="image">

			</div>
			</div>
			<table class="table">
				
				<thead>
				<tr>
					<th class="schedule">Schedule</th>
					<th class="last-service">Last Service</th>
					<th class="next-service">Next Service</th>
				</tr>
				</thead>
				<tbody>
				<tr>
				
					<td class="schedule">
						<span class="table-label">Schedule</span>
						<div></div><div>Thursday every other week</div>					</td>

					<td class="last-service">
						<span class="table-label">Last Service</span>
						17/12/2020					</td>

					<td class="next-service">
						<span class="table-label">Next Service</span>
						04/01/2021					</td>
				</tr>
				<tr>
					<td colspan="3">
													
						<div class="task-state state-not-completed">
							<span class="indicator"></span>
							<p><strong>Last collection: Not completed.</strong></p>						</div>
					</td>
				</tr>
				<tr>
					<td colspan="3" class="event-actions">
						<div class="btn-wrapper"><a class="btn btn-link" href="/praise/2104/property/100020406685/service/547/g/537">Container Request</a></div><div class="other-wrapper"><div class="btn-wrapper"><a class="btn btn-link" href="#">Feedback</a></div><ul class="other-dropdown"><li"><a class="btn btn-link" href="/praise/2162/property/100020406685/service/537">Crew Behaviour</a></li><li"><a class="btn btn-link" href="/praise/2159/property/100020406685/service/537">Damage to 3rd Party Vehicle</a></li><li"><a class="btn btn-link" href="/praise/2163/property/100020406685/service/537">Damage to Property</a></li><li"><a class="btn btn-link" href="/praise/2148/property/100020406685/service/537">General Enquiry</a></li><li"><a class="btn btn-link" href="/praise/2186/property/100020406685/service/537">Wrongful Removal</a></li><li><a class="btn btn-link" href="/praise/2105/property/100020406685/service/547/g/537">Failure to Deliver Bin</a></li></ul></div>					</td>
				</tr>
								<tr>
					<td colspan="3">
						Please note that missed collections can only be reported within 2 working days of your scheduled collection day.
					</td>
				</tr>
								</tbody>
			</table>
		</div>
	</div>
		<div class="service-wrapper service-id-545 task-id-3226">
		<h3 class="service-name">Green Garden Waste (Subscription)					</h3>
		<div class="service-content">
			<div class="image-wrapper">
			<div class="image">

			</div>
			</div>
			<table class="table">
				
				<thead>
				<tr>
					<th class="schedule">Schedule</th>
					<th class="last-service">Last Service</th>
					<th class="next-service">Next Service</th>
				</tr>
				</thead>
				<tbody>
				<tr>
				
					<td class="schedule">
						<span class="table-label">Schedule</span>
						<div></div><div>Wednesday every 4th week</div>					</td>

					<td class="last-service">
						<span class="table-label">Last Service</span>
						09/12/2020					</td>

					<td class="next-service">
						<span class="table-label">Next Service</span>
						20/01/2021					</td>
				</tr>
				<tr>
					<td colspan="3">
													
						<div class="task-state state-completed">
							<span class="indicator"></span>
							<p><strong>Last collection: Your road was completed on 09/12/2020 at 09:27.</strong></p>						</div>
					</td>
				</tr>
				<tr>
					<td colspan="3" class="event-actions">
						<div class="other-wrapper"><div class="btn-wrapper"><a class="btn btn-link" href="#">Feedback</a></div><ul class="other-dropdown"><li"><a class="btn btn-link" href="/praise/2162/property/100020406685/service/545">Crew Behaviour</a></li><li"><a class="btn btn-link" href="/praise/2159/property/100020406685/service/545">Damage to 3rd Party Vehicle</a></li><li"><a class="btn btn-link" href="/praise/2163/property/100020406685/service/545">Damage to Property</a></li><li"><a class="btn btn-link" href="/praise/2148/property/100020406685/service/545">General Enquiry</a></li><li"><a class="btn btn-link" href="/praise/2186/property/100020406685/service/545">Wrongful Removal</a></li></ul></div>					</td>
				</tr>
								<tr>
					<td colspan="3">
						Please note that missed collections can only be reported within 2 working days of your scheduled collection day.
					</td>
				</tr>
								</tbody>
			</table>
		</div>
	</div>
		<div class="service-wrapper service-id-542 task-id-3223">
		<h3 class="service-name">Food waste					</h3>
		<div class="service-content">
			<div class="image-wrapper">
			<div class="image">

			</div>
			</div>
			<table class="table">
				
				<thead>
				<tr>
					<th class="schedule">Schedule</th>
					<th class="last-service">Last Service</th>
					<th class="next-service">Next Service</th>
				</tr>
				</thead>
				<tbody>
				<tr>
				
					<td class="schedule">
						<span class="table-label">Schedule</span>
						<div>Thursday every week</div><div></div>					</td>

					<td class="last-service">
						<span class="table-label">Last Service</span>
						24/12/2020					</td>

					<td class="next-service">
						<span class="table-label">Next Service</span>
						04/01/2021					</td>
				</tr>
				<tr>
					<td colspan="3">
													
						<div class="task-state state-not-completed">
							<span class="indicator"></span>
							<p><strong>Last collection: Not completed.</strong></p>						</div>
					</td>
				</tr>
				<tr>
					<td colspan="3" class="event-actions">
						<div class="btn-wrapper"><a class="btn btn-link" href="/praise/2104/property/100020406685/service/547/g/542">Container Request</a></div><div class="other-wrapper"><div class="btn-wrapper"><a class="btn btn-link" href="#">Feedback</a></div><ul class="other-dropdown"><li"><a class="btn btn-link" href="/praise/2162/property/100020406685/service/542">Crew Behaviour</a></li><li"><a class="btn btn-link" href="/praise/2159/property/100020406685/service/542">Damage to 3rd Party Vehicle</a></li><li"><a class="btn btn-link" href="/praise/2163/property/100020406685/service/542">Damage to Property</a></li><li"><a class="btn btn-link" href="/praise/2148/property/100020406685/service/542">General Enquiry</a></li><li"><a class="btn btn-link" href="/praise/2186/property/100020406685/service/542">Wrongful Removal</a></li><li><a class="btn btn-link" href="/praise/2105/property/100020406685/service/547/g/542">Failure to Deliver Bin</a></li></ul></div>					</td>
				</tr>
								<tr>
					<td colspan="3">
						Please note that missed collections can only be reported within 2 working days of your scheduled collection day.
					</td>
				</tr>
								</tbody>
			</table>
		</div>
	</div>
		<div class="service-wrapper service-id-535 task-id-3216">
		<h3 class="service-name">Plastic, glass and tins					</h3>
		<div class="service-content">
			<div class="image-wrapper">
			<div class="image">

			</div>
			</div>
			<table class="table">
				
				<thead>
				<tr>
					<th class="schedule">Schedule</th>
					<th class="last-service">Last Service</th>
					<th class="next-service">Next Service</th>
				</tr>
				</thead>
				<tbody>
				<tr>
				
					<td class="schedule">
						<span class="table-label">Schedule</span>
						<div></div><div>Thursday every other week</div>					</td>

					<td class="last-service">
						<span class="table-label">Last Service</span>
						24/12/2020					</td>

					<td class="next-service">
						<span class="table-label">Next Service</span>
						09/01/2021					</td>
				</tr>
				<tr>
					<td colspan="3">
													
						<div class="task-state state-completed">
							<span class="indicator"></span>
							<p><strong>Last collection: Your road was completed on 24/12/2020 at 08:49.</strong></p>						</div>
					</td>
				</tr>
				<tr>
					<td colspan="3" class="event-actions">
						<div class="btn-wrapper"><a class="btn btn-link" href="/praise/2104/property/100020406685/service/547/g/535">Container Request</a></div><div class="other-wrapper"><div class="btn-wrapper"><a class="btn btn-link" href="#">Feedback</a></div><ul class="other-dropdown"><li"><a class="btn btn-link" href="/praise/2162/property/100020406685/service/535">Crew Behaviour</a></li><li"><a class="btn btn-link" href="/praise/2159/property/100020406685/service/535">Damage to 3rd Party Vehicle</a></li><li"><a class="btn btn-link" href="/praise/2163/property/100020406685/service/535">Damage to Property</a></li><li"><a class="btn btn-link" href="/praise/2148/property/100020406685/service/535">General Enquiry</a></li><li"><a class="btn btn-link" href="/praise/2186/property/100020406685/service/535">Wrongful Removal</a></li><li><a class="btn btn-link" href="/praise/2105/property/100020406685/service/547/g/535">Failure to Deliver Bin</a></li></ul></div>					</td>
				</tr>
								<tr>
					<td colspan="3">
						Please note that missed collections can only be reported within 2 working days of your scheduled collection day.
					</td>
				</tr>
								</tbody>
			</table>
		</div>
	</div>

	<div class="service-wrapper">
	   <h3 class="service-name">
	      Batteries, small electrical items and textiles
	   </h3>
	   <div class="service-content">
	      <div class="image-wrapper" style="width: 41%">
	         <div class="image" style="padding: 20px;">
	            <img style="width: 66px; margin-right: 10px;" src="/cl/img/services/battery_bag.png">
	            <img style="width: 88px;" src="/cl/img/services/weee_bag.png">
	            <img style="width: 78px;" src="/cl/img/services/textile_bag.png">
	         </div>
	      </div>
	      <table class="table" style="width: 59%; margin: 7% 0 0 41%;">
	         <tbody>
	            <tr>
	               <td colspan="3">
	                  <p>Please only present one small bag each week as we cannot handle high volumes. If we do not collect your items, bring them back inside and present them for collection next week.</p>
	               </td>
	            </tr>
				<tr>
					<td colspan="3" class="event-actions">
						<div class="btn-wrapper"><a class="btn btn-link" href='/praise/2219/property/100020406685/service/535'>Missed Collection</a></div>
					</td>
				</tr>
	         </tbody>
	      </table>
	   </div>
	</div>
         </div>
               </div>
   </div>
</div>
<script>
   (function ($) {
   	$('.other-wrapper > .btn-wrapper').find('.btn').on('click', function () {
   		$(this).closest('.other-wrapper').toggleClass('active');
   		return false;
   	})
   })(jQuery);
</script></div>
<footer class="container-fluid">
    <script>
    SETTINGS = {};
    SETTINGS["MESSAGE_AREA"] = "message-area";  </script>

  
</footer>
</body>
</html>`)
