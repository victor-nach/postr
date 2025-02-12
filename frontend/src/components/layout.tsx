import React from "react";

export default function Layout({ title, children }: { title: string, children: React.ReactNode }) {
  return (
    <div className="min-h-screen flex justify-center">
      <div className="w-full max-w-[856px] px-8 pt-[130px]">
        <h1 className="text-black font-medium text-[60px] sm:text-[40px] md:text-[50px] text-left">{title}</h1>
        {children}
      </div>
    </div>
  );
}
