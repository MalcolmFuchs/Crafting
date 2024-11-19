import Link from "next/link";

export default function Home() {
  return (
  <div className="grid">
    <div>
      <Link href="/time">
        <button>Zeiterfassung</button>
      </Link>
    </div>
    <div>
      <button>Maps</button>
    </div>
  </div>
  );
}
