export function clickOutside(node: Node) {
	const handleClick = (event: Event) => {
		if (node && !node.contains(event.target as Node) && !event.defaultPrevented) {
			node.dispatchEvent(new CustomEvent("outclick"));
		}
	};

	document.addEventListener("click", handleClick, true);

	return {
		destroy() {
			document.removeEventListener("click", handleClick, true);
		}
	};
}
