local _M = {}

function _M.conversion(value)
	local path = ""
	for i=1,#value do
		local tmp = string.sub(value,i,i)
		if tmp=='\\' then
			path = path..'/'
		else
			path = path..tmp
		end
	end
	return path
end

return _M