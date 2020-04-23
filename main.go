package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/job-hunter-toolkit/job-hunter-toolkit/jobpostings"
	"github.com/spf13/cobra"
)

// ls *.go | grep -v 'helper' | grep -v '_test' | awk -F '.' '{print "\""$1"\""","}'
var sourcesList = []string{
	"0x",
	"15five",
	"21st_century_fox",
	"2u",
	"3m",
	"abacus",
	"accion_systems",
	"adjust",
	"adobe",
	"advanced_disposal",
	"aes",
	"affirm",
	"afresh",
	"air_map",
	"airbnb",
	"airtame",
	"akron_childrens_hospital",
	"alterian",
	"alto",
	"amazee",
	"amenity_analytics",
	"amex",
	"amyris",
	"anchor_free",
	"anchorage",
	"apptio",
	"aptible",
	"aquabyte",
	"arbor",
	"arena_net",
	"astranis",
	"atrium",
	"auth0",
	"auto_zone",
	"aware3",
	"axon",
	"azerion",
	"backer_kit",
	"ball_metalpack",
	"bank_of_america",
	"bazaar_voice",
	"beamly",
	"belle_tire",
	"bench",
	"better_lesson",
	"better_up",
	"big_health",
	"bigscreen",
	"binance",
	"bird",
	"bishop_fox",
	"bitnami",
	"bittrex",
	"bjc",
	"blackfynn",
	"blameless",
	"blockstack",
	"blue_owl",
	"bny_mellon",
	"bombas",
	"boom_supersonic",
	"borg_warner",
	"bose",
	"box_lunch",
	"braintree",
	"brave",
	"bright_bytes",
	"brightwheel",
	"britecore",
	"build_zoom",
	"buildium",
	"bustle",
	"buzzfeed",
	"cae",
	"caliber_collision",
	"calm",
	"cambly",
	"canonical",
	"capital_one",
	"carbon_black",
	"carousell",
	"carta",
	"casper",
	"catalytic",
	"cb_insights",
	"celgene",
	"change",
	"chartio",
	"cheddar",
	"chewy",
	"chimp",
	"chronicled",
	"cigna",
	"cisco_meraki",
	"citrix",
	"city_and_county_of_denver",
	"clerky",
	"click_time",
	"close",
	"cloud_bees",
	"cloudflare",
	"cloudreach",
	"cobalt_robotics",
	"coca_cola",
	"cockroach_labs",
	"codacy",
	"code_ocean",
	"codecademy",
	"coffee_meets_bagel",
	"coinbase",
	"colossal",
	"comcast",
	"common",
	"compass",
	"compass_group",
	"conductor_technologies",
	"conrell_university",
	"consensys",
	"contrast_security",
	"convoy",
	"core_scientific",
	"corelight",
	"coupa",
	"coursera",
	"credit_sesame",
	"crowd_strike",
	"cruise",
	"crux",
	"curalate",
	"dahmakan",
	"darkstore",
	"datica",
	"datto",
	"dazn",
	"degreed",
	"deliveroo",
	"dell",
	"digital_ocean",
	"divvy_homes",
	"dnb",
	"docker",
	"domio",
	"door_dash",
	"dots",
	"dr_chrono",
	"drop",
	"duo_lingo",
	"duo_security",
	"early_warning",
	"eaton",
	"ebury",
	"eden",
	"elastic",
	"embark",
	"epic_games",
	"ericsson",
	"erm",
	"essential",
	"eventbrite",
	"evercommerce",
	"evernote",
	"expedia",
	"expensify",
	"farmers_buisness_network",
	"farmers_insurance",
	"fastly",
	"fat_llama",
	"fedex",
	"fetch_rev",
	"fico",
	"fictiv",
	"finger_food_studios",
	"first",
	"fond",
	"forward",
	"fossa",
	"front",
	"fubo_tv",
	"game_changer",
	"gates_foundation",
	"geckoboard",
	"general_assembly",
	"genospace",
	"get_all",
	"github",
	"gitlab",
	"gizmodo",
	"glow",
	"go_daddy",
	"goat",
	"gojek",
	"gov_predict",
	"gradient",
	"gradle",
	"grammarly",
	"greenpeace",
	"gridspace",
	"grimm",
	"guerrilla",
	"gun_io",
	"gusto",
	"hacker_one",
	"hashicorp",
	"hazel_analytics",
	"hello_alfred",
	"hello_fresh",
	"hello_sign",
	"heycar",
	"hifyre",
	"hopsy",
	"imageworks",
	"impossible_foods",
	"in_vision",
	"inboxlab",
	"influx_db",
	"input",
	"insightsoftware",
	"instabase",
	"instacart",
	"instructure",
	"instrumental_ai",
	"inventables",
	"invitae",
	"iota",
	"iris_automation",
	"iron_ox",
	"ironclad",
	"issuu",
	"jda",
	"jellyfish",
	"jet",
	"job_posting",
	"joor",
	"journera",
	"journy",
	"kantar",
	"kard",
	"karius",
	"kayak",
	"khan_academy",
	"kik",
	"kinnek",
	"kite",
	"klarna",
	"koddi",
	"kohls",
	"kraken",
	"launch_darkly",
	"lending_tree",
	"lever",
	"light_step",
	"lighthouse_studios",
	"lilt",
	"lime",
	"limeade",
	"linux_foundation",
	"live_nation",
	"log_dna",
	"logitech",
	"looker",
	"lucid",
	"lyft",
	"lyric",
	"magic_leap",
	"major_league_baseball",
	"make_school",
	"make_space",
	"managed_by_q",
	"mapbox",
	"marriott",
	"mastercard",
	"mattermost",
	"maven_clinic",
	"mavens",
	"mc_donalds",
	"mckesson",
	"measurabl",
	"med_men",
	"medium",
	"message_bird",
	"metal_toad",
	"mighty_networks",
	"modernize",
	"modsy",
	"mongodb",
	"motorola_solutions",
	"mount_sinai",
	"movable_ink",
	"mux",
	"narvar",
	"nash",
	"nationwide",
	"netlify",
	"neuralink",
	"new_engen",
	"new_york_times",
	"newfront_insurance",
	"nexient",
	"nextdoor",
	"niantic",
	"nova",
	"npm",
	"ns1",
	"nurx",
	"nutanix",
	"nvidia",
	"ny_media",
	"nylas",
	"ochsner",
	"okta",
	"omada_health",
	"omaze",
	"one_login",
	"open_ai",
	"open_cosmos",
	"open_fin",
	"open_gov",
	"open_market",
	"open_phone",
	"opendoor",
	"optiv",
	"origin",
	"outschool",
	"pachyderm",
	"pae",
	"pager_duty",
	"palantir",
	"palo_alto_networks",
	"pantheon",
	"paperspace",
	"path_ai",
	"pathlight",
	"patreon",
	"pax",
	"paytm",
	"people_ai",
	"pepsico",
	"persist_iq",
	"petal",
	"pfizer",
	"pinterest",
	"pioneer_square_labs",
	"placepass",
	"plan_grid",
	"platformsh",
	"plume",
	"policygenius",
	"poll_everywhere",
	"popdog",
	"postmates",
	"praetorian",
	"predictive_index",
	"procore",
	"projekt202",
	"proofpoint",
	"protocol",
	"psi_kick",
	"pushpay",
	"pwc",
	"qualtrics",
	"qualys",
	"quartzy",
	"quidd",
	"quora",
	"rainforest_qa",
	"rapid7",
	"recurly",
	"reddit",
	"reify_health",
	"relativity",
	"relayr",
	"remarkably",
	"remitly",
	"remix",
	"replate",
	"research_gate",
	"returntocorp",
	"ring",
	"riot_games",
	"robinhood",
	"roblox",
	"roll_pay",
	"rolls_royce",
	"rooster_teeth",
	"rover",
	"rubrik",
	"rush",
	"salesforce",
	"salt_lending",
	"salt_stack",
	"samsara",
	"samsung",
	"sanofi",
	"sauce_labs",
	"scale_ai",
	"scandit",
	"scoop",
	"screen_cloud",
	"scribd",
	"second_measure",
	"second_spectrum",
	"security_trails",
	"securonix",
	"semaphore_solutions",
	"sensor_tower",
	"sentry",
	"shape_scale",
	"shoes",
	"shop_keep",
	"shopify",
	"sift",
	"signal",
	"simpli_safe",
	"skillshare",
	"skip",
	"skydio",
	"slack",
	"smarking",
	"smartsheet",
	"snap",
	"snap_raise",
	"snap_travel",
	"snapdocs",
	"snowflake",
	"sony_play_station",
	"source_d",
	"sourceress",
	"spaceflight_industries",
	"spacex",
	"sparkswap",
	"spot_hero",
	"spotfront",
	"springboard",
	"sprout_social",
	"squarespace",
	"stack_adapt",
	"starcity",
	"starship",
	"starsky_robotics",
	"state_steet",
	"stauer",
	"strava",
	"streamlabs",
	"stripe",
	"sumo_logic",
	"survey_gizmo",
	"survey_monkey",
	"swat",
	"swiss_borg",
	"symantec",
	"sysco",
	"sysdig",
	"t1cg",
	"tableau",
	"tablexi",
	"taco_bell",
	"tailored_brands",
	"tala",
	"taplytics",
	"teamable",
	"telaria",
	"the_athletic",
	"the_trade_desk",
	"thumbtack",
	"thunder_token",
	"tilting_point",
	"time",
	"time_inc",
	"top_hat",
	"toptal",
	"trans_lifeline",
	"transcend",
	"transparent_systems",
	"trilogy_education",
	"trip_advisor",
	"triple_byte",
	"triplemint",
	"ttt_studios",
	"tuft_and_needle",
	"tune",
	"tutela",
	"twilio",
	"twitch",
	"tyro",
	"uber",
	"uberether",
	"udacity",
	"uhaul",
	"uniregistry",
	"unisys",
	"unity3d",
	"universal_music_group",
	"university_of_chicago",
	"university_of_virgina",
	"upstart",
	"upwork",
	"user_leap",
	"venafi",
	"vend",
	"venmo",
	"verge_genomics",
	"verifone",
	"verishop",
	"veritas",
	"verizon_media",
	"vewd",
	"vice",
	"vidyard",
	"visa",
	"voleon",
	"voodoo",
	"vox_media",
	"voxter",
	"wealthsimple",
	"web3",
	"wellframe",
	"wgames",
	"wheels",
	"whisper",
	"whole_foods",
	"wikimedia",
	"wizeline",
	"wonder",
	"workday",
	"xylem",
	"zendesk",
	"zenefits",
	"zentail",
	"zeplin",
	"zero_cater",
	"zerofox",
	"zeus",
	"zipline",
	"zipwhip",
	"zume",
	"zyris",
}

func main() {
	cleanup := func() {
		os.Exit(0)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			cleanup()
		}
	}()

	var (
		cmdJobPostingsPrintJSON bool
		cmdJobPostingsPrintCSV  bool
	)

	var cmdJobPostings = &cobra.Command{
		Use:   "job-postings [flags]",
		Short: "Find job postings from various companies",
		Run: func(cmd *cobra.Command, args []string) {
			var printer func(j *jobpostings.JobPosting)

			if cmdJobPostingsPrintJSON {
				printer = func(j *jobpostings.JobPosting) {
					data, err := json.Marshal(j)
					if err == nil {
						fmt.Println(string(data))
					}
				}
			} else if cmdJobPostingsPrintCSV {
				printerWrapper := func() func(j *jobpostings.JobPosting) {
					w := csv.NewWriter(os.Stdout)

					return func(j *jobpostings.JobPosting) {
						record := []string{j.Company, j.Title, j.Location, j.URL}
						if err := w.Write(record); err != nil {
							panic(err)
						}
					}
				}
				printer = printerWrapper()
			} else {
				printer = func(j *jobpostings.JobPosting) {
					fmt.Println("company:", j.Company, "title:", j.Title, "location:", j.Location, "url:", j.URL)
				}
			}

			for jobPosting := range jobpostings.GetAllJobPostings(context.Background()) {
				printer(jobPosting)
			}
		},
	}
	cmdJobPostings.Flags().BoolVar(&cmdJobPostingsPrintJSON, "json", false, "output as newline separated JSON")
	cmdJobPostings.Flags().BoolVar(&cmdJobPostingsPrintCSV, "csv", false, "output as CSV with no header (title, location, url)")

	var (
		cmdJobSourcesPrintJSON bool
		cmdJobSourcesPrintCSV  bool
	)

	var cmdJobSources = &cobra.Command{
		Use:   "job-sources [flags]",
		Short: "List the various companies available from job-postings",
		Run: func(cmd *cobra.Command, args []string) {
			var printer func(s []string)

			if cmdJobSourcesPrintJSON {
				printer = func(s []string) {
					data, err := json.Marshal(s)
					if err == nil {
						fmt.Println(string(data))
					}
				}
			} else if cmdJobSourcesPrintCSV {
				printer = func(s []string) {
					fmt.Println(strings.Join(s, ","))
				}
			} else {
				printer = func(s []string) {
					for _, source := range s {
						fmt.Println(source)
					}
				}
			}

			printer(sourcesList)
		},
	}
	cmdJobSources.Flags().BoolVar(&cmdJobSourcesPrintJSON, "json", false, "output as newline separated JSON")
	cmdJobSources.Flags().BoolVar(&cmdJobSourcesPrintCSV, "csv", false, "output as CSV with no header (source1, sourc2, ...)")

	var rootCmd = &cobra.Command{Use: "job-hunter-toolkit"}
	rootCmd.AddCommand(cmdJobPostings)
	rootCmd.AddCommand(cmdJobSources)
	rootCmd.Execute()
}
