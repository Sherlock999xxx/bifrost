import { Gauge } from "lucide-react";
import ContactUsView from "../views/contactUsView";

export default function UserLimitsView() {
	return (
		<div className="h-full w-full">
			<ContactUsView
				className="mx-auto min-h-[80vh]"
				icon={<Gauge className="h-[5.5rem] w-[5.5rem]" strokeWidth={1} />}
				title="Unlock user-level budgets and rate limits"
				description="Set spending limits and rate controls per user to manage costs and prevent abuse. This feature is part of the Bifrost enterprise license."
				readmeLink="https://docs.getbifrost.ai/enterprise/user-governance"
			/>
		</div>
	);
}
