export function timePassed(date) {
  const now = new Date();
  const pastDate = new Date(date);
  const diff = now - pastDate; // Difference in milliseconds

  const seconds = Math.floor(diff / 1000);
  const minutes = Math.floor(seconds / 60);
  const hours = Math.floor(minutes / 60);
  const days = Math.floor(hours / 24);

  if (seconds < 60) {
    return `${seconds} seconds ago`;
  } else if (minutes < 60) {
    return `${minutes} minutes ago`;
  } else if (hours < 24) {
    return `${hours} hours ago`;
  } else {
    return `${days} days ago`;
  }
}

// Example usage
const pastDate = "2025-01-14T12:00:00"; // Replace with your date
console.log(timePassed(pastDate));
