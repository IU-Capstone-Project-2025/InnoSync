import { ChevronDown, Bell } from "lucide-react";

const Navbar = () => {
  return (
    <nav className="w-full h-[100px] bg-white relative">
      <div className="flex items-center justify-between w-full h-full px-4 lg:px-10">
        {/* Logo */}
        <div className="flex items-center">
          <h1 className="text-3xl lg:text-[48px] font-bold font-montserrat leading-normal">
            <span className="text-inno-green">Inno</span>
            <span className="text-black">Sync</span>
          </h1>
        </div>

        {/* Navigation Links - Hidden on mobile, shown on large screens */}
        <div className="hidden lg:flex items-center gap-5">
          <div className="flex items-center gap-1 cursor-pointer">
            <span className="text-black font-bold font-montserrat text-2xl">
              Find Project
            </span>
            <ChevronDown className="w-4 h-4 text-black" strokeWidth={3} />
          </div>
          <div className="flex items-center gap-1 cursor-pointer">
            <span className="text-black font-bold font-montserrat text-2xl">
              Find Talent
            </span>
            <ChevronDown className="w-4 h-4 text-black" strokeWidth={3} />
          </div>
          <div className="flex items-center gap-1 cursor-pointer">
            <span className="text-black font-bold font-montserrat text-2xl">
              About Us
            </span>
            <ChevronDown className="w-4 h-4 text-black" strokeWidth={3} />
          </div>
        </div>

        {/* User Section */}
        <div className="flex items-center gap-1">
          {/* Notifications */}
          <div className="relative mr-1">
            <Bell className="w-6 h-6 lg:w-7 lg:h-7 text-inno-green fill-inno-green" />
          </div>

          {/* Role Selector - Hidden text on small screens */}
          <div className="flex items-center justify-center px-2 lg:px-[15px] h-[29px] cursor-pointer">
            <span className="hidden sm:block text-inno-green font-bold font-montserrat text-lg lg:text-2xl">
              Talent
            </span>
            <div className="ml-1">
              <svg
                width="21"
                height="21"
                viewBox="0 0 21 21"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  fillRule="evenodd"
                  clipRule="evenodd"
                  d="M10.2549 2.33005C5.74274 2.33005 2.08493 5.98785 2.08493 10.5C2.08493 15.0121 5.74274 18.6699 10.2549 18.6699C14.767 18.6699 18.4248 15.0121 18.4248 10.5C18.4248 5.98785 14.767 2.33005 10.2549 2.33005ZM0.04245 10.5C0.0424502 4.85982 4.61471 0.287567 10.2549 0.287567C15.895 0.287567 20.4673 4.85982 20.4673 10.5C20.4673 16.1401 15.895 20.7124 10.2549 20.7124C4.61471 20.7124 0.0424497 16.1401 0.04245 10.5ZM6.46902 8.75662C6.86784 8.3578 7.51445 8.3578 7.91327 8.75662L10.2549 11.0982L12.5965 8.75662C12.9953 8.3578 13.6419 8.3578 14.0407 8.75662C14.4395 9.15544 14.4395 9.80205 14.0407 10.2009L10.977 13.2646C10.7855 13.4561 10.5257 13.5637 10.2549 13.5637C9.98402 13.5637 9.72426 13.4561 9.53274 13.2646L6.46902 10.2009C6.0702 9.80205 6.0702 9.15544 6.46902 8.75662Z"
                  fill="#298217"
                />
              </svg>
            </div>
          </div>

          {/* User Dropdown */}
          <div className="flex items-center justify-center px-2 lg:px-[15px] h-[29px] cursor-pointer">
            <span className="text-inno-green font-bold font-montserrat text-lg lg:text-2xl">
              A.Alimi
            </span>
            <div className="ml-1 rotate-90">
              <svg
                width="21"
                height="21"
                viewBox="0 0 21 21"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  fillRule="evenodd"
                  clipRule="evenodd"
                  d="M2.08496 10.5C2.08496 15.0121 5.74277 18.6699 10.2549 18.6699C14.767 18.6699 18.4248 15.0121 18.4248 10.5C18.4248 5.98788 14.767 2.33008 10.2549 2.33008C5.74277 2.33008 2.08496 5.98788 2.08496 10.5ZM10.2549 20.7124C4.61474 20.7124 0.0424802 16.1402 0.0424805 10.5C0.0424807 4.85985 4.61474 0.287594 10.2549 0.287594C15.8951 0.287594 20.4673 4.85985 20.4673 10.5C20.4673 16.1402 15.8951 20.7124 10.2549 20.7124ZM8.51153 14.2859C8.11271 13.887 8.11271 13.2404 8.51153 12.8416L10.8531 10.5L8.51153 8.15841C8.11271 7.7596 8.11271 7.11298 8.51153 6.71416C8.91035 6.31534 9.55696 6.31534 9.95578 6.71416L13.0195 9.77789C13.211 9.96941 13.3186 10.2292 13.3186 10.5C13.3186 10.7709 13.211 11.0306 13.0195 11.2221L9.95578 14.2859C9.55696 14.6847 8.91035 14.6847 8.51153 14.2859Z"
                  fill="#298217"
                />
              </svg>
            </div>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
